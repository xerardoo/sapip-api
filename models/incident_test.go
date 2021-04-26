package models_test

import (
	"fmt"
	"github.com/xerardoo/sapip/models"
	"testing"

	"encoding/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	_ "github.com/xerardoo/sapip/tester"
)

func TestIncident(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Incident Test Suite")
}

var _ = Describe("Incident", func() {

	incident := new(models.Incident)
	// incident.Type =
	incident.Description = "se hizo alto preventivo"
	// incident.UserID = 1
	incident.LocationID = 1

	// Add new
	incidentAdded, err := incident.Add()
	It("cant be added", func() {
		Expect(err).NotTo(HaveOccurred())
		Expect(incidentAdded.Description).Should(Equal("se hizo alto preventivo"))
	})

	var count int64
	err = models.DB.Model(models.Incident{}).Count(&count).Error
	It("cant be added", func() {
		Expect(err).NotTo(HaveOccurred())
		fmt.Println("COUNT", count)
	})

	location, err := incidentAdded.GetLocation()
	It("cant be added", func() {
		Expect(err).NotTo(HaveOccurred())
		res, _ := json.Marshal(location)
		fmt.Println("ROWS", string(res))
	})
	user, err := incidentAdded.GetUser()
	It("cant be added", func() {
		Expect(err).NotTo(HaveOccurred())
		Expect(user.FirstName).Should(Equal("Juan"))
	})

	var incidents []models.Incident
	err = models.DB.Find(&incidents).Error
	It("cant be added", func() {
		Expect(err).NotTo(HaveOccurred())
	})

	// incident := new(models.Incident2)
	// incident.Date = "2020-05-05"
	// incident.Time = "13:24:45"
	// incident.Description = "se hizo alto preventivo"
	//
	// err := models.DB.Create(&incident).Error
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
})
