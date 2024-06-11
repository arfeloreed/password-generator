package internal

import (
	"encoding/json"
	"fmt"
	"os"
)

type Data struct {
	Email    string
	Password string
}

var message string = "success"

const filename = "passwords.json"

func CreateJson(email string, password string) (string, error) {
	// check if the json file exists, if not create one
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		data := []Data{
			{Email: email, Password: password},
		}
		// convert data to json
		jsonData, err := json.MarshalIndent(data, "", "\t")
		if err != nil {
			Screen()
			fmt.Println("Can't convert data to json\n", err)
			return "", err
		}
		// write jsondata to json file
		err = os.WriteFile(filename, jsonData, 0644)
		if err != nil {
			Screen()
			fmt.Println("Can't write json data\n", err)
			return "", err
		}
	} else {
		// if file exists, add new data
		data := []Data{}
		// read json file
		jsonData, err := os.ReadFile(filename)
		if err != nil {
			Screen()
			fmt.Println("Can't read json file\n", err)
			return "", err
		}
		// convert json to data
		err = json.Unmarshal(jsonData, &data)
		if err != nil {
			Screen()
			fmt.Println("Can't convert json to data\n", err)
			return "", err
		}
		// add new data
		data = append(data, Data{Email: email, Password: password})
		// convert the updated data to json
		updatedJsonData, err := json.MarshalIndent(data, "", "\t")
		if err != nil {
			Screen()
			fmt.Println("Can't convert data to json\n", err)
			return "", err
		}
		// write jsondata to json file
		err = os.WriteFile(filename, updatedJsonData, 0644)
		if err != nil {
			Screen()
			fmt.Println("Can't write json data\n", err)
			return "", err
		}
	}

	return message, nil
}

func SearchJson(email string) (Data, error) {
	// check if json file exists, then read json file
	_, err := os.Stat(filename)
	if err == nil {
		jsonData, err := os.ReadFile(filename)
		if err != nil {
			Screen()
			fmt.Println("Can't read json file\n", err)
			return Data{}, err
		}
		// convert json to data
		var data []Data
		err = json.Unmarshal(jsonData, &data)
		if err != nil {
			Screen()
			fmt.Println("Can't convert json to data\n", err)
			return Data{}, err
		}
		// search for email
		for _, d := range data {
			if d.Email == email {
				return d, nil
			}
		}
	}
	return Data{}, nil
}

func EditJson(email string, password string) (string, error) {
	// check if json file exists, then read json file
	_, err := os.Stat(filename)
	if err == nil {
		// if file exists, add new data
		data := []Data{}
		// read json file
		jsonData, err := os.ReadFile(filename)
		if err != nil {
			Screen()
			fmt.Println("Can't read json file\n", err)
			return "", err
		}
		// convert json to data
		err = json.Unmarshal(jsonData, &data)
		if err != nil {
			Screen()
			fmt.Println("Can't convert json to data\n", err)
			return "", err
		}
		// search for email and update password
		for i, d := range data {
			fmt.Println(d.Email)
			if d.Email == email {
				data[i].Password = password
				// convert the updated data to json
				updatedJsonData, err := json.MarshalIndent(data, "", "\t")
				if err != nil {
					Screen()
					fmt.Println("Can't convert data to json\n", err)
					return "", err
				}
				// write jsondata to json file
				err = os.WriteFile(filename, updatedJsonData, 0644)
				if err != nil {
					Screen()
					fmt.Println("Can't write json data\n", err)
					return "", err
				}
				return message, nil
			}
		}
	}
	return "", nil
}

func DeleteEmail(email string) (string, error) {
	// check if json file exists, then read json file
	_, err := os.Stat(filename)
	if err == nil {
		// if file exists, edit the data
		data := []Data{}
		// read json file
		jsonData, err := os.ReadFile(filename)
		if err != nil {
			Screen()
			fmt.Println("Can't read json file\n", err)
			return "", err
		}
		// convert json to data
		err = json.Unmarshal(jsonData, &data)
		if err != nil {
			Screen()
			fmt.Println("Can't convert json to data\n", err)
			return "", err
		}
		// search for email and delete
		for i, d := range data {
			if d.Email == email {
				data = append(data[:i], data[i+1:]...)
				// convert the updated data to json
				updatedJsonData, err := json.MarshalIndent(data, "", "\t")
				if err != nil {
					Screen()
					fmt.Println("Can't convert data to json\n", err)
					return "", err
				}
				// write jsondata to json file
				err = os.WriteFile(filename, updatedJsonData, 0644)
				if err != nil {
					Screen()
					fmt.Println("Can't write json data\n", err)
					return "", err
				}
				return message, nil
			}
		}
	}
	return "", nil
}
