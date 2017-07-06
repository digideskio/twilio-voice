package main_test

import (
	"os"

	. "github.com/BTBurke/twilio-voice"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config", func() {
	var cfg *Config

	BeforeEach(func() {
		cfg = &Config{
			MailgunPublicKey:  "abc123",
			MailgunSecretKey:  "pancakes",
			MailgunDomain:     "example.com",
			ForwardingNumber:  "+15555555",
			NotificationEmail: "voicemail@example.com",
		}
	})

	Describe(".Validate", func() {
		It("returns empty slice when valid", func() {
			Expect(cfg.Validate()).To(BeEmpty())
		})

		It("sets default VoicemailScript when VoicemailFile does not exist", func() {
			cfg.VoicemailFile = "/path/to/a/nonexistent/file.mp3"
			cfg.VoicemailScript = ""

			Expect(cfg.Validate()).To(BeEmpty())
			Expect(cfg.VoicemailScript).To(Equal("Please leave a message"))
		})

		// FIXME(ivy): blank VoicemailFile succeeds
		XIt("sets default VoicemailScript when VoicemailFile unspecificed", func() {
			cfg.VoicemailFile = ""
			cfg.VoicemailScript = ""

			Expect(cfg.Validate()).To(BeEmpty())
			Expect(cfg.VoicemailScript).To(Equal("Please leave a message"))
		})

		It("uses specified VoicemailFile when exists", func() {
			wd, err := os.Getwd()
			Expect(err).NotTo(HaveOccurred())

			cfg.VoicemailFile = "README.md"

			Expect(cfg.Validate()).To(BeEmpty())
			Expect(cfg.ServeDirectory).To(Equal(wd + "/"))
			Expect(cfg.VoiceFileName).To(Equal("README.md"))
		})

		It("returns error when missing MailgunPublicKey", func() {
			cfg.MailgunPublicKey = ""

			errs := cfg.Validate()
			Expect(len(errs)).To(Equal(1))
			Expect(errs[0]).To(MatchError(
				MatchRegexp("set.*MAILGUN_PUBLIC_KEY"),
			))
		})

		It("returns error when missing MailgunSecretKey", func() {
			cfg.MailgunSecretKey = ""

			errs := cfg.Validate()
			Expect(len(errs)).To(Equal(1))
			Expect(errs[0]).To(MatchError(
				MatchRegexp("set.*MAILGUN_SECRET_KEY"),
			))
		})

		It("returns error when missing MailgunDomain", func() {
			cfg.MailgunDomain = ""

			errs := cfg.Validate()
			Expect(len(errs)).To(Equal(1))
			Expect(errs[0]).To(MatchError(
				MatchRegexp("set.*MAILGUN_DOMAIN"),
			))
		})

		It("returns error when missing NotificationEmail", func() {
			cfg.NotificationEmail = ""

			errs := cfg.Validate()
			Expect(len(errs)).To(Equal(1))
			Expect(errs[0]).To(MatchError(
				MatchRegexp("set.*NOTIFICATION_EMAIL"),
			))
		})

		It("returns error when missing ForwardingNumber", func() {
			cfg.ForwardingNumber = ""

			errs := cfg.Validate()
			Expect(len(errs)).To(Equal(1))
			Expect(errs[0]).To(MatchError(
				MatchRegexp("set.*FORWARDING_NUMBER"),
			))
		})
	})
})
