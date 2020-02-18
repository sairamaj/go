terraform {
  # Generic modules should pin a RANGE of Terraform versions.
  # Live modules should pin a SPECIFIC Terraform version.
  required_version = ">=0.12, <0.13"
}

provider "sample" {
  path = "c:\\temp"
}

resource sample_item test_file {
    name = "this_is_an_item"
    filename = "test.txt"
    content = "somecontent"
}
