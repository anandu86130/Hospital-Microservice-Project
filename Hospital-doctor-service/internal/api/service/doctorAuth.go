package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/anandu86130/Hospital-doctor-service/config"
	generateotp "github.com/anandu86130/Hospital-doctor-service/internal/generateOTP"
	"github.com/anandu86130/Hospital-doctor-service/internal/model"
	pb "github.com/anandu86130/Hospital-doctor-service/internal/pbD"
	"github.com/anandu86130/Hospital-doctor-service/internal/token"
	"github.com/anandu86130/Hospital-doctor-service/internal/utility"
	"golang.org/x/crypto/bcrypt"
)

func (u *DoctorService) LoginService(doctorpb *pb.Login) (*pb.Response, error) {
	// Find the doctor by email
	user, err := u.Repo.FindDoctorByEmail(doctorpb.Email)
	if err != nil {
		// Return error if user not found
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "user not found",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, errors.New("user not found")
	}

	// Compare the hashed password with the provided password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(doctorpb.Password))
	if err != nil {
		// If password comparison fails
		log.Printf("Password comparison failed: %v", err)
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "password comparison failed",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, errors.New("password comparison failed")
	}

	// Check if the user is blocked
	if user.IsBlocked {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "user is blocked by Admin",
		}, errors.New("you are blocked by Admin")
	}

	if !user.IsVerified {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "doctor is not verified by Admin",
		}, errors.New("you are not verified by Admin")
	}
	// Continue with token generation...
	token, _ := token.GenerateToken(config.LoadConfig().SECRETKEY, user.Email, uint(user.ID))

	// Return successful login response with token
	return &pb.Response{
		Status:  pb.Response_OK,
		Message: "Successfully logged in",
		Payload: &pb.Response_Data{Data: token},
	}, nil
}

func (u *DoctorService) SignupService(doctorpb *pb.Signup) (*pb.Response, error) {
	existingUser, err := u.Repo.FindDoctorByEmail(doctorpb.Email)
	if err == nil && existingUser != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "email already exists",
		}, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(doctorpb.Password), bcrypt.DefaultCost)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "error in hashing password",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, errors.New("unable to hashpassword")
	}

	if len(doctorpb.Password) <= 3 {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "password length not met",
			Payload: &pb.Response_Error{Error: "password should be minimum 4 digits or numbers"},
		}, errors.New("please enter min 4 digits or numbers")
	}

	if doctorpb.Specialization != "dermatology" && doctorpb.Specialization != "neurology" && doctorpb.Specialization != "cardiology" && doctorpb.Specialization != "medicine" && doctorpb.Specialization != "oncology" && doctorpb.Specialization != "anesthesiology" && doctorpb.Specialization != "gynacology" && doctorpb.Specialization != "pediatrics" && doctorpb.Specialization != "psychiatrist" && doctorpb.Specialization != "ophthalmology" && doctorpb.Specialization != "surgeon" && doctorpb.Specialization != "radiology" {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "please enter a valid specialization",
			Payload: &pb.Response_Error{Error: ""},
		}, errors.New("incorrect specialization")
	}

	// Generate a new OTP
	otp := generateotp.GenerateOTP(6)
	err = utility.SendOTPByEmail(doctorpb.Email, otp)
	if err != nil{
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "failed to send otp via email",
			Payload: &pb.Response_Error{Error: ""},
		}, errors.New("failed to send otp via email")	
	}

	// Check if an OTP already exists for the given email
	existingOTP, err := u.Repo.FindOTPByEmail(doctorpb.Email)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "error in sending otp to email using smtp",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, errors.New("failed to send OTP")
	}

	if existingOTP != nil {
		// Update the existing OTP entry
		existingOTP.Otp = otp
		existingOTP.Password = string(hashedPassword)
		err = u.Repo.UpdateOTP(existingOTP)
		if err != nil {
			return &pb.Response{
				Status:  pb.Response_ERROR,
				Message: "error in updating OTP",
				Payload: &pb.Response_Error{Error: err.Error()},
			}, errors.New("failed to update OTP")
		}
	} else {
		// Create a new OTP entry
		create := model.OTP{
			Email:    doctorpb.Email,
			Otp:      otp,
			Password: string(hashedPassword),
		}

		err = u.Repo.CreateOTP(&create)
		if err != nil {
			return &pb.Response{
				Status:  pb.Response_ERROR,
				Message: "error in creating OTP",
				Payload: &pb.Response_Error{Error: err.Error()},
			}, errors.New("failed to create OTP")
		}
	}

	doctorpb.Password = string(hashedPassword)

	uniqueLicense, err := utility.GenerateUniqueLicense(10) // Specify desired length
	if err != nil {
		log.Fatalf("Failed to generate license number: %v", err)
	}
	// Store the OTP in Redis with an expiration time
	storedata := &model.Doctor{Name: doctorpb.Name, Email: doctorpb.Email, Password: doctorpb.Password, Specialization: doctorpb.Specialization, YearsOfExperience: doctorpb.YearsOfExperience, LicenceNumber: uniqueLicense, Fees: doctorpb.Fees}
	userData, err := json.Marshal(&storedata)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "error in marshaling user data",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, errors.New("error while marshaling data")
	}

	key := fmt.Sprintf("user:%s:%s", doctorpb.Email, otp)
	err = u.Redis.SetDataInRedis(key, userData, time.Minute*10)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "error in setting data in redis",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, errors.New("error while setting data in redis")
	}

	return &pb.Response{
		Status:  pb.Response_OK,
		Message: "OTP send successfully, Please verify OTP",
	}, nil
}

func (u *DoctorService) VerifyOTP(doctorpb *pb.OTP) (*pb.Response, error) {
	log.Printf("Verifying OTP for email: %s", doctorpb.Email)

	// Verify OTP from the database
	err := u.Repo.VerifyOTPcheck(doctorpb.Email, doctorpb.Otp)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "failed to verify OTP",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, errors.New("failed to verify OTP")
	}

	// Generate Redis key for the OTP
	key := fmt.Sprintf("user:%s:%s", doctorpb.Email, doctorpb.Otp)
	log.Printf("Retrieving OTP data from Redis with key: %s", key)

	// Retrieve OTP data from Redis
	userData, err := u.Redis.GetFromRedis(key)
	if err != nil {
		log.Printf("Error recieving data from Redis: %v", err)
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "error when recieving data from redis",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, errors.New("error when recieving data from redis")
	}

	if userData == "" {
		log.Printf("No data found in Redis for key: %s", key)
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "user data not found in redis",
		}, errors.New("no data in redis")
	}

	// Unmarshal the data retrieved from Redis
	var verifyOTPs model.Doctor
	err = json.Unmarshal([]byte(userData), &verifyOTPs)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "failed to unmarshal user data",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, errors.New("error in unmarshaling user data")
	}

	// Validate that the password is a valid bcrypt hash
	if len(verifyOTPs.Password) != 60 {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "invalid password hash stored in redis",
		}, errors.New("invalid password hash stored in redis")
	}

	// Create a new user using the unmarshaled data
	user := &model.Doctor{
		Name:              verifyOTPs.Name,
		Email:             verifyOTPs.Email,
		Password:          verifyOTPs.Password, // Already hashed password
		Specialization:    verifyOTPs.Specialization,
		YearsOfExperience: verifyOTPs.YearsOfExperience,
		LicenceNumber:     verifyOTPs.LicenceNumber,
		Fees:              verifyOTPs.Fees,
	}

	// Check if the user already exists
	exists, err := u.Repo.DoctorExists(user.Email)
	if err != nil {
		return &pb.Response{
			Status:  pb.Response_ERROR,
			Message: "failed to check user exists",
			Payload: &pb.Response_Error{Error: err.Error()},
		}, errors.New("failed to check user exists")
	}

	if exists {
		// Update the existing user
		err = u.Repo.UpdateDoctor(user)
		if err != nil {
			return &pb.Response{
				Status:  pb.Response_ERROR,
				Message: "failed to update user",
				Payload: &pb.Response_Error{Error: err.Error()},
			}, errors.New("failed to update user details")
		}
	} else {
		// Create a new user
		err = u.Repo.CreateDoctor(user)
		if err != nil {
			return &pb.Response{
				Status:  pb.Response_ERROR,
				Message: "failed to create user",
				Payload: &pb.Response_Error{Error: err.Error()},
			}, errors.New("error in creating user")
		}
	}

	log.Printf("User successfully created and token generated for email: %s", doctorpb.Email)
	return &pb.Response{
		Status:  pb.Response_OK,
		Message: "OTP verified successfully, login to continue",
	}, nil
}
