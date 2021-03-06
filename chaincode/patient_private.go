/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/chaincode/shim/ext/cid"
	sc "github.com/hyperledger/fabric/protos/peer"
)

//SmartContract struct
type SmartContract struct {
}

//Doctor struct
type Doctor struct {
	DoctorID      string `json:"dockorId"`
	DoctorName    string `json:"dockorName"`
	DoctorSurname string `json:"dockorSurname"`
}

//Patient struct
type Patient struct {
	PatientID      string   `json:"patientID"`
	PatientName    string   `json:"patientName"`
	PatientSurname string   `json:"patientSurname"`
	PatientFileURL string   `json:"patientFileUrl"`
	PatientDoctors []Doctor `json:"patientDoctors"`
}

//Init method
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

//Invoke method
/*Access Contol:
* Only dev1 can update the ledger
* Only dev2 can read the ledger
 */
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	function, args := APIstub.GetFunctionAndParameters()
	if function == "addPatient" {
		err := cid.AssertAttributeValue(APIstub, "dev1", "true")
		if err != nil {
			return shim.Error("Unauthorised access dectected fagster")
		}
		return s.addPatient(APIstub, args)
	} else if function == "getPatient" {
		err := cid.AssertAttributeValue(APIstub, "dev2", "true")
		if err != nil {
			return shim.Error("Unauthorised access dectected fagster")
		}
		return s.getPatient(APIstub, args)
	} else if function == "updatePatient" {
		err := cid.AssertAttributeValue(APIstub, "dev1", "true")
		if err != nil {
			return shim.Error("Unauthorised access dectected fagster")
		}
		return s.updatePatient(APIstub, args)
	} else if function == "addDoctor" {
		err := cid.AssertAttributeValue(APIstub, "dev1", "true")
		if err != nil {
			return shim.Error(err.Error())
		}
		return s.addDoctor(APIstub, args)

	} else if function == "getPatientHistory" {
		err := cid.AssertAttributeValue(APIstub, "dev2", "true")
		if err != nil {
			return shim.Error(err.Error())
		}
		return s.getHistoryForPatient(APIstub, args)
	} else if function == "addPatientPrivate" {
		err := cid.AssertAttributeValue(APIstub, "dev1", "true")
		if err != nil {
			return shim.Error(err.Error())
		}
		return s.addPatientPrivate(APIstub, args)
	} else if function == "getPatientPrivate" {
		err := cid.AssertAttributeValue(APIstub, "dev2", "true")
		if err != nil {
			return shim.Error(err.Error())
		}
		return s.getPatientPrivate(APIstub, args)
	} else if function == "initLedger" {
		return s.initLedger(APIstub)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	doctor := Doctor{DoctorID: "pp", DoctorName: "Paul", DoctorSurname: "Potts"}
	doctorAsBytes, _ := json.Marshal(doctor)
	APIstub.PutState("doctorkey", doctorAsBytes)

	patient := Patient{PatientID: "testid", PatientName: "testname", PatientSurname: "testsurname", PatientFileURL: "testUrl"}
	patient.PatientDoctors = append(patient.PatientDoctors, doctor)
	patientAsBytes, _ := json.Marshal(patient)
	APIstub.PutState("testkey", patientAsBytes)
	return shim.Success(nil)
}

//Add normal patient
func (s *SmartContract) addPatient(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}
	patient := Patient{PatientID: args[1], PatientName: args[2], PatientSurname: args[3], PatientFileURL: args[4]}

	patientAsBytes, _ := json.Marshal(patient)
	APIstub.PutState(args[0], patientAsBytes)

	return shim.Success(nil)
}

//Update normal patient
func (s *SmartContract) updatePatient(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}
	patientAsBytes, _ := APIstub.GetState(args[0])
	patient := Patient{}

	json.Unmarshal(patientAsBytes, &patient)
	patient.PatientName = args[1]
	patient.PatientSurname = args[2]
	patient.PatientFileURL = args[3]

	patientAsBytes, _ = json.Marshal(patient)
	APIstub.PutState(args[0], patientAsBytes)

	return shim.Success(nil)
}

//Normal get patient
func (s *SmartContract) getPatient(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	patientAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(patientAsBytes)
}

//Add Patient privately
func (s *SmartContract) addPatientPrivate(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	transMap, err := APIstub.GetTransient()

	if _, ok := transMap["patient_details"]; !ok {
		return shim.Error("marble must be a key in the transient map")
	}

	var patientTransient Patient
	err = json.Unmarshal(transMap["patient_details"], &patientTransient)
	if err != nil {
		return shim.Error("Failed to decode JSON of: " + string(transMap["patient_details"]))
	}

	patient := Patient{
		PatientID:      patientTransient.PatientID,
		PatientName:    patientTransient.PatientName,
		PatientSurname: patientTransient.PatientSurname,
		PatientFileURL: patientTransient.PatientFileURL,
	}

	//patient := Patient{PatientID: args[1], PatientName: args[2], PatientSurname: args[3], PatientFileURL: args[4]}

	patientAsBytes, _ := json.Marshal(patient)
	APIstub.PutPrivateData("private_patient_collection", patientTransient.PatientID, patientAsBytes)

	return shim.Success(nil)
}

//Get Patient privately
func (s *SmartContract) getPatientPrivate(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	patientAsBytes, _ := APIstub.GetPrivateData("private_patient_collection", args[0])
	return shim.Success(patientAsBytes)
}

//Add a Doctor
func (s *SmartContract) addDoctor(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	doctor := Doctor{DoctorID: args[1], DoctorName: args[2], DoctorSurname: args[3]}
	doctorAsBytes, _ := json.Marshal(doctor)
	APIstub.PutState(args[0], doctorAsBytes)
	return shim.Success(nil)
}

//Get normal Patient history
func (s *SmartContract) getHistoryForPatient(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	patientName := args[0]

	fmt.Printf("- start getHistoryForMarble: %s\n", patientName)

	resultsIterator, err := stub.GetHistoryForKey(patientName)

	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing historic values for the marble
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(response.TxId)
		buffer.WriteString("\"")
		buffer.WriteString(", \"IsDelete\":")
		buffer.WriteString("\"")
		buffer.WriteString(strconv.FormatBool(response.IsDelete))
		buffer.WriteString("\"")

		buffer.WriteString(", \"Value\":")
		// if it was a delete operation on given key, then we need to set the
		//corresponding value null. Else, we will write the response.Value
		//as-is (as the Value itself a JSON marble)
		if response.IsDelete {
			buffer.WriteString("null")
		} else {
			buffer.WriteString(string(response.Value))
		}

		buffer.WriteString(", \"Timestamp\":")
		buffer.WriteString("\"")
		buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
		buffer.WriteString("\"")

		buffer.WriteString(", \"IsDelete\":")
		buffer.WriteString("\"")
		buffer.WriteString(strconv.FormatBool(response.IsDelete))
		buffer.WriteString("\"")

		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getHistoryPatient returning:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())

}

func main() {
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
