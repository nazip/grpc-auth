package v1

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate minimock -i API -o ./mocks/ -s "_minimock.go"
