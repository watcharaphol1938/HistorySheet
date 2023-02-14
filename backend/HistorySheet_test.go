package main

import (
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
	"gorm.io/gorm"
)

type HistorySheet struct {
	gorm.Model
	Name   string    `valid:"required~Name should not be blank"`
	Age    int       `valid:"int, range(0|100)~Age should not be negative integer"`
	Weight float32   `valid:"float, range(0|100)~Weight should not be negative float"`
	Url    string    `valid:"required~Url should not be blank, url~Url should be match"`
	Mobile string    `valid:"required~Mobile should not be blank, matches(^[0]{1}[689]{1}[0-9]{8})~Mobile should be match"`
	Date   time.Time `valid:"present"`
}

func TestNameMustNotBeBlank(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	t.Run("Check Name must be not blank", func(t *testing.T) {
		h := HistorySheet{
			Name:   "",
			Age:    23,
			Weight: 56.23,
			Url:    "www.google.com",
			Mobile: "0635946211",
			Date:   time.Now(),
		}
		ok, err := govalidator.ValidateStruct(h)
		g.Expect(ok).ToNot(gomega.BeTrue())
		g.Expect(err).NotTo(gomega.BeNil())
		g.Expect(err.Error()).To(gomega.Equal("Name should not be blank"))
	})
}
func TestAgeMustNotBeNegativeInteger(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	t.Run("Check Age must be not negative integer", func(t *testing.T) {
		h := HistorySheet{
			Name:   "Teerasil",
			Age:    -23,
			Weight: 56.23,
			Url:    "www.google.com",
			Mobile: "0635946211",
			Date:   time.Now(),
		}
		ok, err := govalidator.ValidateStruct(h)
		g.Expect(ok).ToNot(gomega.BeTrue())
		g.Expect(err).NotTo(gomega.BeNil())
		g.Expect(err.Error()).To(gomega.Equal("Age should not be negative integer"))
	})
}
func TestWeightMustNotBeNegativeFloat(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	t.Run("Check Weight must be not negative float", func(t *testing.T) {
		h := HistorySheet{
			Name:   "Teerasil",
			Age:    23,
			Weight: -56.23,
			Url:    "www.google.com",
			Mobile: "0635946211",
			Date:   time.Now(),
		}
		ok, err := govalidator.ValidateStruct(h)
		g.Expect(ok).ToNot(gomega.BeTrue())
		g.Expect(err).NotTo(gomega.BeNil())
		g.Expect(err.Error()).To(gomega.Equal("Weight should not be negative float"))
	})
}
func TestUrlMustNotBeBlank(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	t.Run("Check Url must not be blank", func(t *testing.T) {
		h := HistorySheet{
			Name:   "Teerasil",
			Age:    23,
			Weight: 56.23,
			Url:    "",
			Mobile: "0635946211",
			Date:   time.Now(),
		}
		ok, err := govalidator.ValidateStruct(h)
		g.Expect(ok).ToNot(gomega.BeTrue())
		g.Expect(err).NotTo(gomega.BeNil())
		g.Expect(err.Error()).To(gomega.Equal("Url should not be blank"))
	})
}
func TestUrlMustBeMatch(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	t.Run("Check Url must be match", func(t *testing.T) {
		h := HistorySheet{
			Name:   "Teerasil",
			Age:    23,
			Weight: 56.23,
			Url:    "wwwgooglecom",
			Mobile: "0635946211",
			Date:   time.Now(),
		}
		ok, err := govalidator.ValidateStruct(h)
		g.Expect(ok).ToNot(gomega.BeTrue())
		g.Expect(err).NotTo(gomega.BeNil())
		g.Expect(err.Error()).To(gomega.Equal("Url should be match"))
	})
}
func TestMobileMustNotBeBlank(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	t.Run("Check Mobile must not be match", func(t *testing.T) {
		h := HistorySheet{
			Name:   "Teerasil",
			Age:    23,
			Weight: 56.23,
			Url:    "www.google.com",
			Mobile: "",
			Date:   time.Now(),
		}
		ok, err := govalidator.ValidateStruct(h)
		g.Expect(ok).ToNot(gomega.BeTrue())
		g.Expect(err).NotTo(gomega.BeNil())
		g.Expect(err.Error()).To(gomega.Equal("Mobile should not be blank"))
	})
}
func TestMobileMustBeMatch(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	t.Run("Check Mobile must be match", func(t *testing.T) {
		h := HistorySheet{
			Name:   "Teerasil",
			Age:    23,
			Weight: 56.23,
			Url:    "www.google.com",
			Mobile: "0321659946",
			Date:   time.Now(),
		}
		ok, err := govalidator.ValidateStruct(h)
		g.Expect(ok).ToNot(gomega.BeTrue())
		g.Expect(err).NotTo(gomega.BeNil())
		g.Expect(err.Error()).To(gomega.Equal("Mobile should be match"))
	})
}
func TestDateMustNotBeFuture(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	// m := "03 Feb 2066"
	tm, _ := time.Parse("03 Feb 2000", "01 Feb 2066")
	t.Run("Check Date must not be past", func(t *testing.T) {
		h := HistorySheet{
			Name:   "Teerasil",
			Age:    23,
			Weight: 56.23,
			Url:    "www.google.com",
			Mobile: "0921659946",
			Date:   tm,
		}
		ok, err := govalidator.ValidateStruct(h)
		g.Expect(ok).ToNot(gomega.BeTrue())
		g.Expect(err).NotTo(gomega.BeNil())
		g.Expect(err.Error()).To(gomega.Equal("Date should not be future"))
	})
}
func init() {
	govalidator.CustomTypeTagMap.Set("past", func(i interface{}, context interface{}) bool {
		t := i.(time.Time)
		return t.Before(time.Now())
	})

	govalidator.CustomTypeTagMap.Set("future", func(i interface{}, context interface{}) bool {
		t := i.(time.Time)
		return t.After(time.Now())
	})

	govalidator.CustomTypeTagMap.Set("present", func(i interface{}, o interface{}) bool {
		t := i.(time.Time)
		// ปัจจุบัน บวกลบไม่เกิน 12 ชั่วโมง
		return t.After(time.Now().Add(time.Hour*-12)) && t.Before(time.Now().Add(time.Hour*12))
	})
}
