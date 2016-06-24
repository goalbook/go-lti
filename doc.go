/*
	Implementation of LTI 1.0 Protocol in Go.

	Currently there are a few major missing features such as the implementation of Tool Consumer, common cartridge parsing, proper nonce and timestamp verification.


	Tool Provider Usage

	Tool Provider wraps a http.Request and provides easy access to LTI headers, Validation of LTI requests and creation of return URLS.

	LTI headers are deserialized into a LTIStdHeaders struct which can be accessed from the LTIHeaders attribute of an LTIToolProvider.

		func Handler(w http.ResponseWriter, r *http.Request) {

		    // Create a new LTIToolProvider
			ltiRequest, err := lti.NewLTIToolProvider(r)

		    // Validate LTI request
		    valid, err := ltiRequest.ValidateRequest(secretKey, false, false)

			if valid == true {

		        // Access some LTI Header
		    	fmt.Println(ltiRequest.LTIHeaders.LISPersonFamilyName)

		    } else {

		    	// Redirect to return URL
		        returnUrl, _ = ltiRequest.CreateReturnURL()
		        http.Redirect(w, r, returnUrl.String(), http.StatusMovedPermanently)

		    }
		}
*/
package lti
