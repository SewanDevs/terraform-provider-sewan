# Terraform resources
* component of terraform provider

# Cheat sheet
* return nils = operation succeeded without error

## CREATE call

## READ call

## UPDATE call
* sync local state with actual upstream state
* handle case where resource doesn't exists : no error return

## DELETE call
* delete the ressource if it exists only
* do not return an error if the resource doesn't exists : terraform standard to prevent terraform instance of breaking
