// Copyright © 2019 VMware, INC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package purgedb

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	client "github.com/edgexfoundry/edgex-cli/pkg"
	"github.com/edgexfoundry/go-mod-core-contracts/models"
)

// NewCommand returns the purgedb command
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "purgedb",
		Short: "Purges entire EdgeX Database. [USE WITH CAUTION]",
		Long: `Purge DB

USE WITH CAUTION. The effect of this command is irreversible.

The purgedb command purges the entire Database. It performs the same action as the
clean_mongo.js developer script. Unlike the clean_mongo.js, this command purges the 
database using API calls only. clean_mongo.js accesses the DB directly, which might 
always be possible using the CLI.
`,
		Run: func(cmd *cobra.Command, args []string) {

			//////////////////////////////////////////////////////
			// COMMAND
			//////////////////////////////////////////////////////
			type commandList struct {
				list []models.Command
			}

			commandData := client.GetAllItems("command")

			commands := commandList{}
			commanderrjson := json.Unmarshal(commandData, &commands.list)
			if commanderrjson != nil {
				fmt.Println(commanderrjson)
			}

			fmt.Println("COMMAND--------------------------")
			numberItems := len(commands.list)
			for _, object := range commands.list {
				fmt.Println(object.Id)
				fmt.Println(object.Name)

				// call delete function here
				client.DeleteItem(object.Id, "command")
			}
			fmt.Println("Removed ", numberItems, " commands.")

			//////////////////////////////////////////////////////
			// DEVICE
			//////////////////////////////////////////////////////
			type deviceList struct {
				list []models.Device
			}

			deviceData := client.GetAllItems("device")

			devices := deviceList{}
			deviceerrjson := json.Unmarshal(deviceData, &devices.list)
			if deviceerrjson != nil {
				fmt.Println(deviceerrjson)
			}

			fmt.Println("DEVICE----------------")
			for _, object := range devices.list {
				fmt.Println(object.Id)
				fmt.Println(object.Name)
				// call delete function here
			}

			//////////////////////////////////////////////////////
			// DP
			//////////////////////////////////////////////////////
			type deviceProfileList struct {
				list []models.DeviceProfile
			}

			DeviceProfileData := client.GetAllItems("deviceprofile")

			deviceprofiles := deviceProfileList{}

			deviceprofileerrjson := json.Unmarshal(DeviceProfileData, &deviceprofiles.list)
			if deviceprofileerrjson != nil {
				fmt.Println(deviceprofileerrjson)
			}

			fmt.Println("-DP---------------------------")
			for _, object := range deviceprofiles.list {
				fmt.Println(object.Id)
				fmt.Println(object.Name)

				// call delete function here
			}

			//////////////////////////////////////////////////////
			// DR
			//////////////////////////////////////////////////////

			type deviceReportList struct {
				list []models.DeviceReport
			}

			deviceReportData := client.GetAllItems("devicerepost")

			devicereports := deviceReportList{}

			devicereporterrjson := json.Unmarshal(deviceReportData, &devicereports.list)
			if devicereporterrjson != nil {
				fmt.Println(devicereporterrjson)
			}

			fmt.Println("-----DR----------------------------------")
			for _, object := range devicereports.list {
				fmt.Println(object.Id)
				fmt.Println(object.Name)

				// call delete function here
			}

			//////////////////////////////////////////////////////
			// DS
			//////////////////////////////////////////////////////
			type deviceServiceList struct {
				list []models.DeviceService
			}

			deviceServiceData := client.GetAllItems("deviceservice")

			deviceservices := deviceServiceList{}

			deviceserviceerrjson := json.Unmarshal(deviceServiceData, &deviceservices.list)
			if deviceserviceerrjson != nil {
				fmt.Println(deviceserviceerrjson)
			}

			fmt.Println("--DS---------------------------------------")
			for _, object := range deviceservices.list {
				fmt.Println(object.Id)
				fmt.Println(object.Name)

				// call delete function here
			}

			//////////////////////////////////////////////////////
			// ADDRESSABLES
			//////////////////////////////////////////////////////

			type addressableList struct {
				list []models.Addressable
			}

			// Calling GetAllItems function, which
			// makes API call to get all items of given typ
			data := client.GetAllItems("addressable")

			// unmarshalling the json response
			list := addressableList{}
			errjson := json.Unmarshal(data, &list.list)
			if errjson != nil {
				fmt.Println(errjson)
			}

			// Looping over the list of items and calling
			// DeleteItem for each
			fmt.Println("ADDRESSABLES----------------------")
			// numberItems := len(list.list)
			for _, addr := range list.list {
				fmt.Println(addr.Id)
				fmt.Println(addr.Name)

				// call delete function here
				// client.DeleteItem(addr.Id, "addressable")
			}
			fmt.Println(numberItems)

			//////////////////////////////////////////////////////
			// Provision watchers
			//////////////////////////////////////////////////////

			type provisionWatcherList struct {
				list []models.ProvisionWatcher
			}

			provisionWatcherData := client.GetAllItems("provisionwatcher")

			provisionwatchers := provisionWatcherList{}

			provisionwatchererrjson := json.Unmarshal(provisionWatcherData, &provisionwatchers.list)
			if provisionwatchererrjson != nil {
				fmt.Println(provisionwatchererrjson)
			}

			fmt.Println("-Provision watchers----------------------------")
			for _, object := range provisionwatchers.list {
				fmt.Println(object.Id)
				fmt.Println(object.Name)

				// call delete function here
			}

			//////////////////////////////////////////////////////
			// Schedule
			//////////////////////////////////////////////////////

			type scheduleList struct {
				list []models.ProvisionWatcher
			}

			scheduleData := client.GetAllItems("schedule")

			schedules := scheduleList{}

			schedulerjson := json.Unmarshal(scheduleData, &schedules.list)
			if schedulerjson != nil {
				fmt.Println(schedulerjson)
			}

			fmt.Println("---Schedule--------------------------------")
			for _, object := range schedules.list {
				fmt.Println(object.Id)
				fmt.Println(object.Name)

				// call delete function here
			}

			// TODO: check why not working:

			// type scheduleEventList struct {
			// 	addr []models
			// }

			// DONE:
			// Meta data:
			// addressable
			// command
			// device
			// deviceProfile
			// deviceReport
			// deviceService
			// provisionWatcher
			// schedule
			// scheduleEvent

			// TODO
			// coredata:
			// valueDescriptor
			// reading: http://localhost:48080/api/v1/event/scrub Might have to delete one by one
			// event: http://localhost:48080/api/v1/event/scrub

			// logging:
			// logEntry /logs/{start}/{end}

			// notifications:
			// notification http://localhost:48060/api/v1/cleanup
			// subscription might be affected by /cleanup
			// transmission same

			// exportclient:
			// exportConfiguration one by one

		},
	}
	return cmd
}
