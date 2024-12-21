package utils

import (
	"booking-service/config"
	"fmt"
	"log"

	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentintent"
)

// StripeClient holds the Stripe configuration and methods to interact with the Stripe API.
type StripeClient struct {
	apiKey string
	redis  *config.RedisService // Assuming a Redis client is used
}

// NewStripeClient initializes the Stripe client with the API key from the configuration.
func NewStripeClient(cfg config.Config, redis *config.RedisService) *StripeClient {
	stripe.Key = cfg.STRIPEKEY
	if stripe.Key == "" {
		log.Fatal("Stripe secret key is missing from configuration")
	}
	return &StripeClient{
		apiKey: cfg.STRIPEKEY,
		redis:  redis,
	}
}

// CreatePaymentIntent creates a Stripe PaymentIntent and returns its ID and client secret.
func (s *StripeClient) CreatePaymentIntent(amount int64, currency string) (string, string, error) {
	// Validate the minimum allowed amount for the given currency
	if !isValidAmount(amount, currency) {
		return "", "", fmt.Errorf("amount is below the minimum allowed amount for the currency")
	}

	// Create the PaymentIntent with the specified amount and currency
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(amount),
		Currency: stripe.String(currency),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}

	// Create the PaymentIntent
	pi, err := paymentintent.New(params)
	if err != nil {
		log.Printf("Stripe API error: %v", err)
		return "", "", fmt.Errorf("failed to create payment intent: %v", err)
	}

	// Store the payment ID and client secret in Redis for later use
	err = s.storeClientSecretInRedis(pi.ID, pi.ClientSecret)
	if err != nil {
		return "", "", fmt.Errorf("failed to store client secret in Redis: %v", err)
	}

	// Return the PaymentIntent ID and client secret
	return pi.ID, pi.ClientSecret, nil
}

// storeClientSecretInRedis stores the PaymentIntent's client secret in Redis.
func (s *StripeClient) storeClientSecretInRedis(paymentID, clientSecret string) error {
	// err := s.redis.SetDataInRedis(fmt.Sprintf("payment:%s", paymentID), clientSecret, 0)
	err := s.redis.SetDataInRedis(fmt.Sprintf("payment:%s", paymentID), []byte(clientSecret), 0)

	if err != nil {
		log.Printf("Failed to stor e client secret for payment %s in Redis: %v", paymentID, err)
		return err
	}
	return nil
}

// VerifyPaymentStatus retrieves the status of a payment intent by its ID.
func (s *StripeClient) VerifyPaymentStatus(paymentID string) (string, error) {
	stripe.Key = s.apiKey // Use the client's Stripe API key

	intent, err := paymentintent.Get(paymentID, nil)
	if err != nil {
		log.Printf("Failed to retrieve payment intent: %v", err)
		return "", fmt.Errorf("error retrieving payment intent: %v", err)
	}

	// Return the status of the payment intent (e.g., succeeded, failed, etc.)
	return string(intent.Status), nil
}

// isValidAmount checks if the amount is above the minimum allowed for the given currency.
func isValidAmount(amount int64, currency string) bool {
	switch currency {
	case "usd":
		// Minimum for USD is $0.50 (i.e., 50 cents)
		return amount >= 50
	case "inr":
		// Minimum for INR is ₹50 (i.e., 5000 paise)
		return amount >= 5000
	default:
		// For unsupported currencies, return false
		return false
	}
}
