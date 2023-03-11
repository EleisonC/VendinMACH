package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/EleisonC/vending-machine/helpers"
	"github.com/EleisonC/vending-machine/models"
	"github.com/gorilla/mux"
)

func CreateNewProductHn(w http.ResponseWriter, r *http.Request) {
	var product models.ProductModel
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		helpers.VenErrorHandler(w, "Product information error", err)
		return
	}

	usernameST, err := helpers.ExtractClaims(w, r)
	if err != nil {
		helpers.VenErrorHandler(w, "Claims Issue", err)
		return
	}
	if usernameST["role"].(string) != "seller" {
		helpers.VenErrorHandler(w, "Not enough right to perform this", errors.New("restricted"))
		return
	}

	product.SellerID = usernameST["userId"].(string)

	if err := validate.Struct(&product); err != nil {
		helpers.VenErrorHandler(w, "Product information error", err)
		return
	}

	if err := models.CreateNewProduct(&product); err != nil {
		helpers.VenErrorHandler(w, "product not created", err)
		return
	}

	postRes := models.PosMessageRes{
		Message: "Product created",
	}

	res, err := json.Marshal(postRes)
	if err != nil {
		helpers.VenErrorHandler(w, "Somthing Happened. But User Is Create", err)
		return
	}

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetAllProductsHn(w http.ResponseWriter, r *http.Request) {
	var productsList []models.ProductModel

	err := models.GetAllProducts(&productsList)
	if err != nil {
		helpers.VenErrorHandler(w, "Failed to get Items", err)
		return
	}

	res, err := json.Marshal(productsList)
	if err != nil {
		helpers.VenErrorHandler(w, "Issue Getting User", err)
		return
	}

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateProductHn(w http.ResponseWriter, r *http.Request) {
	var editProduct models.ProductModelUp
	var currentProduct models.ProductModel
	params := mux.Vars(r)
	productId := params["productId"]

	usernameST, err := helpers.ExtractClaims(w, r)
	if err != nil {
		helpers.VenErrorHandler(w, "Claims Issue", err)
		return
	}
	if usernameST["role"].(string) != "seller" && currentProduct.SellerID == usernameST["userId"].(string) {
		helpers.VenErrorHandler(w, "Not enough right to perform this", errors.New("restricted"))
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&editProduct); err != nil {
		helpers.VenErrorHandler(w, "Product update failed data error", err)
	}

	if err := models.GetProductById(&currentProduct, productId); err != nil {
		helpers.VenErrorHandler(w, "Product probably does not exist or nor enough rights to update", err)
		return
	}

	if err := validate.Struct(&editProduct); err != nil {
		helpers.VenErrorHandler(w, "Product update failed data error", err)
		return
	}

	if err := models.UpdateProduct(&editProduct, productId, usernameST["userId"].(string)); err != nil {
		helpers.VenErrorHandler(w, "Something Happened During Update", err)
		return
	}

	postRes := models.PosMessageRes{
		Message: "Product updated",
	}

	res, err := json.Marshal(postRes)
	if err != nil {
		helpers.VenErrorHandler(w, "Somthing Happened. But Product Is Updated", err)
		return
	}

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteProductByIdHn(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	productId := params["productId"]

	usernameST, err := helpers.ExtractClaims(w, r)
	if err != nil {
		helpers.VenErrorHandler(w, "Claims Issue", err)
		return
	}
	if usernameST["role"].(string) != "seller" {
		helpers.VenErrorHandler(w, "Not enough right to perform this", errors.New("restricted"))
		return
	}

	err = models.DeleteProductById(productId, usernameST["userId"].(string))
	if err != nil {
		helpers.VenErrorHandler(w, "Not Enough Rights", err)
		return
	}

	postRes := models.PosMessageRes{
		Message: "Product Deleted",
	}

	res, err := json.Marshal(postRes)
	if err != nil {
		helpers.VenErrorHandler(w, "Somthing Happened. But Product Is Updated", err)
		return
	}

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


