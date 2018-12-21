// Copyright 2018 BlueData Software, Inc.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import "encoding/json"

// UnmarshalJSON for SetupPackage handles the unmarshalling of three
// scenarios wrt 'setup_package':
//   1. omitted                 : IsSet==false
//   2. explicitly set to null  : IsSet==true && IsNull==true
//   3. Set to a valid object   : IsSet=true && IsNull==false
func (setupPackage *SetupPackage) UnmarshalJSON(
	data []byte,
) error {

	// The fact that we entered this function means the filed is set otherwise,
	// this field will be false by default.
	setupPackage.IsSet = true

	dataStr := string(data)
	if (len(dataStr) == 0) || (dataStr == "null") {
		// The field value is explicitly set to null
		setupPackage.IsNull = true
		return nil
	}

	if err := json.Unmarshal(data, &setupPackage.PackageURL); err != nil {
		return err
	}
	setupPackage.IsNull = false

	return nil
}

// UnmarshalJSON for Image handles the unmarshalling of three
// scenarios wrt 'image_repo_tag':
//   1. omitted                 : IsSet==false
//   2. explicitly set to null  : IsSet==true && IsNull==true
//   3. Set to a valid object   : IsSet=true && IsNull==false
func (image *Image) UnmarshalJSON(
	data []byte,
) error {

	// The fact that we entered this function means the filed is set otherwise,
	// this field will be false by default.
	image.IsSet = true

	dataStr := string(data)
	if (len(dataStr) == 0) || (dataStr == "null") {
		// The field value is explicitly set to null
		image.IsNull = true
		return nil
	}

	if err := json.Unmarshal(data, &image.RepoTag); err != nil {
		return err
	}
	image.IsNull = false

	return nil
}

// // MarshalJSON is to handle encoding the Image into json.
// func (image Image) MarshalJSON() ([]byte, error) {
// 	if image.IsSet == false {
// 		return json.Marshall(nil)
// 	}
// 	return
// }