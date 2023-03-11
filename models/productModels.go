package models

import (
	"errors"

	"github.com/EleisonC/vending-machine/config"
	_ "github.com/mattn/go-sqlite3"
)




func CreateNewProduct(product *ProductModel) error {
	config.ConnectDB()
	stmt, err := config.DB.Prepare("INSERT INTO productsTable(productname, sellerid, cost, amountavailable) VALUES(?, ?, ?, ?)")
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(product.ProductName, product.SellerID, product.Cost, product.AmountAvailable)
	if err != nil {
		return err
	}
	return nil
}

func UpdateProduct(productUpdate *ProductModelUp, productId string, sellerId string) error {
	config.ConnectDB()
	stmt, err := config.DB.Prepare("UPDATE productTable SET productname=?, cost=?, amountavaiable=? WHERE id=? AND sellerid=?")
	if err != nil {
		return err
	}

	defer stmt.Close()
	res, err := stmt.Exec(productUpdate.ProductName, productUpdate.Cost, productUpdate.AmountAvailable, productId, sellerId)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil || count == 0{
        return errors.New("there was an issue updating the product")
    }

	return nil
}

func GetAllProducts(productModelSlice *[]ProductModel) (error){
	config.ConnectDB()
	var product ProductModel
	rows, err := config.DB.Query("SELECT * FROM productTable")
	if err != nil {
		return err
	}

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&product.Id, &product.ProductName, &product.SellerID, &product.Cost, &product.AmountAvailable)
		if err != nil {
			return err
		}
		*productModelSlice = append(*productModelSlice, product)
	}
	return nil
}

func GetProductById(product *ProductModel, productId string) (error){
	config.ConnectDB()
	rows := config.DB.QueryRow("SELECT * FROM productTable WHERE id=?", productId)
	err := rows.Scan(&product.Id, &product.ProductName, &product.SellerID, &product.Cost, &product.AmountAvailable)
	if err != nil {
		return err
	}
	return nil
}

func UpdateProductAmt(productAmt int, productId string) (error) {
	config.ConnectDB()
	stmt, err := config.DB.Prepare("UPDATE productTable SET amountavailable=? WHERE id=?")
	if err != nil {
		return err
	}

	defer stmt.Close()
	rows, err:= stmt.Exec(productAmt, productId)
	if err != nil {
		return err
	}

	count, err := rows.RowsAffected()
	if err != nil || count == 0 {
		return errors.New("something happened during the amount update")
	}
	return nil
}

func DeleteProductById(productId string, userId string) (error){
	config.ConnectDB()
	stmt, err := config.DB.Prepare("DELETE FROM productTable WHERE id=? AND sellerid=?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(productId, userId)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil || count == 0 {
		return errors.New("something happened during the delete proccess")
	}
	return nil
}


