package model

type car struct {
	ID    int
	Brand string
	Model string
	Year  int
	Price float64
}

var Cars = []car{
	{ID: 1, Brand: "Toyota", Model: "Camry", Year: 2021, Price: 24425.00},
	{ID: 2, Brand: "Honda", Model: "Civic", Year: 2020, Price: 20650.00},
	{ID: 3, Brand: "Ford", Model: "Mustang", Year: 2022, Price: 27155.00},
	{ID: 4, Brand: "Chevrolet", Model: "Malibu", Year: 2021, Price: 22800.00},
	{ID: 5, Brand: "Tesla", Model: "Model 3", Year: 2022, Price: 39999.00},
	{ID: 6, Brand: "BMW", Model: "3 Series", Year: 2021, Price: 41300.00},
	{ID: 7, Brand: "Audi", Model: "A4", Year: 2020, Price: 39900.00},
	{ID: 8, Brand: "Mercedes-Benz", Model: "C-Class", Year: 2022, Price: 41900.00},
	{ID: 9, Brand: "Volkswagen", Model: "Passat", Year: 2021, Price: 23990.00},
	{ID: 10, Brand: "Hyundai", Model: "Sonata", Year: 2020, Price: 23950.00},
	{ID: 11, Brand: "Kia", Model: "Optima", Year: 2021, Price: 23690.00},
	{ID: 12, Brand: "Subaru", Model: "Outback", Year: 2022, Price: 26645.00},
	{ID: 13, Brand: "Mazda", Model: "Mazda6", Year: 2020, Price: 24100.00},
	{ID: 14, Brand: "Nissan", Model: "Altima", Year: 2021, Price: 24650.00},
	{ID: 15, Brand: "Jeep", Model: "Cherokee", Year: 2020, Price: 26990.00},
	{ID: 16, Brand: "Lexus", Model: "ES", Year: 2022, Price: 39900.00},
	{ID: 17, Brand: "Volvo", Model: "S60", Year: 2021, Price: 38650.00},
	{ID: 18, Brand: "Porsche", Model: "911", Year: 2020, Price: 101200.00},
	{ID: 19, Brand: "Jaguar", Model: "XF", Year: 2022, Price: 45100.00},
	{ID: 20, Brand: "Land Rover", Model: "Range Rover", Year: 2021, Price: 92000.00},
	{ID: 21, Brand: "Cadillac", Model: "CT5", Year: 2020, Price: 37000.00},
	{ID: 22, Brand: "Acura", Model: "TLX", Year: 2022, Price: 37700.00},
	{ID: 23, Brand: "Alfa Romeo", Model: "Giulia", Year: 2021, Price: 41145.00},
	{ID: 24, Brand: "Infiniti", Model: "Q50", Year: 2020, Price: 36200.00},
	{ID: 25, Brand: "Genesis", Model: "G70", Year: 2022, Price: 37950.00},
}
