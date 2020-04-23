package purge

import (
	"fmt"
	"github.com/edgexfoundry-holding/edgex-cli/config"
	client "github.com/edgexfoundry-holding/edgex-cli/pkg"
	"github.com/edgexfoundry/go-mod-core-contracts/clients"
	"github.com/edgexfoundry/go-mod-core-contracts/models"
)

type CoreDataCleaner interface {
	Purge()
	cleanEvents()
	cleanReadings()
	cleanValueDescriptors()
}

type coredataCleaner struct {
	baseUrl string
}

// NewCoredataCleaner creates an instance of CoreDataCleaner
func NewCoredataCleaner() CoreDataCleaner {
	fmt.Println("\n * core-data")
	return &coredataCleaner{
		baseUrl: config.Conf.Clients["CoreData"].Url(),
	}
}
func (d *coredataCleaner) Purge() {
	d.cleanEvents()
	d.cleanReadings()
	d.cleanValueDescriptors()
}

func (d *coredataCleaner) cleanReadings() {
	url := d.baseUrl + clients.ApiReadingRoute
	var readings []models.Reading
	err := client.ListHelper(url, readings)
	if err != nil {
		fmt.Println(err)
		return
	}

	var count int
	for _, reading := range readings {
		// call delete function here
		_, err = client.DeleteItem(url + config.PathId + reading.Id)
		if err != nil {
			fmt.Printf("Failed to delete Reading with id %s because of error: %s", reading.Id, err)
		} else {
			count = count +1
		}
	}
	fmt.Printf("Removed %d Reading from %d \n", count, len(readings))
}

func (d *coredataCleaner) cleanValueDescriptors() {
	url := d.baseUrl + clients.ApiValueDescriptorRoute
	var valueDescriptors []models.ValueDescriptor
	err := client.ListHelper(url, valueDescriptors)
	if err != nil {
		fmt.Println(err)
		return
	}

	var count int
	for _, valueDescriptor := range valueDescriptors {
		_, err = client.DeleteItem(url + config.PathId + valueDescriptor.Id)
		if err != nil {
			fmt.Printf("Failed to delete Value Descriptor with id %s because of error: %s", valueDescriptor.Id, err)
		} else {
			count = count +1
		}
	}
	fmt.Printf("Removed %d Value Descriptors from %d \n", count, len(valueDescriptors))
}

func (d *coredataCleaner) cleanEvents() {
	url := d.baseUrl + clients.ApiEventRoute + "/scruball"
	_, err := client.DeleteItem(url)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Print("All Events have been removed \n")
	}
}