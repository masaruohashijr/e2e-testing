Feature: Newman Postman API Testing    

Scenario Outline: Testing newman apis
  Given The env File "cpaas.postman_environment.json"
  And The data file <data_file> and collection folder <collection_folder>
  When Running Newman Collection "ffb6e8669787e0ed8d13"
  Then All test cases must pass

  Examples:
    | data_file | collection_folder |
    |    "add_new_number_data.json" |  "AddNewNumber-Local"  |
    |    "delete_number_data.json" |   "DeleteNumber-Local" |


  # Background:
  #   Given The env File "cpaas.postman_environment.json"

  # Scenario: API Testing for Adding New Number
    
  #   And The data file "AddNewNumber-Local/add_new_number_data.json"
  #   When Running Newman Collection "ffb6e8669787e0ed8d13" and folder "AddNewNumber-Local"
  #   Then All test cases must pass


  # Scenario: API Testing for Deleting Number
  #   And The data file "DeleteNumber-Local/delete_number_data.json"
  #   When Running Newman Collection "ffb6e8669787e0ed8d13" and folder "DeleteNumber-Local"
  #   Then All test cases must pass