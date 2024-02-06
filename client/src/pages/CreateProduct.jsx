import React, { useState } from "react";
import Cookies from "js-cookie";
import { Footer, Navbar } from "../components";
import Select from "react-select";
import { useParams } from "react-router-dom";

const CreateProduct = () => {
  const { id } = useParams();

  const [selectedImage, setSelectedImage] = useState(null);
  const [title, setTitle] = useState("");
  const [price, setPrice] = useState("");
  const [description, setDescription] = useState("");
  const [category, setCategory] = useState("");
  const [image, setImage] = useState(null);

  const handleImageChange = (event) => {
    const file = event.target.files[0];
    console.log(event.target.files[0]);

    if (file) {
      const reader = new FileReader();

      reader.onload = (e) => {
        setSelectedImage(e.target.result);
      };

      reader.readAsDataURL(file);
    }
    setImage(event.target.files[0]);
  };

  const handleSelectChange = (selectedOption) => {
    setCategory(selectedOption.value);
  };

  // async function handleSubmit() {
  const handleSubmit = async (event) => {
    try {
      if (
        title == "" ||
        price == "" ||
        description == "" ||
        category == "" ||
        image == ""
      ) {
        throw "Error : Invalid value";
      }

      const priceParse = parseFloat(price);

      const values = {
        title: title,
        price: priceParse,
        description: description,
        category: category,
        image: image,
      };

      let response = await fetch(`http://127.0.0.1:8080/product/`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          authorization: `Bearer ${Cookies.get("jwt")}`,
        },
        body: JSON.stringify(values),
      });
      window.location.replace("http://localhost:3000/producttable");
      if (response.status === 200) {
        console.log("Upload successful!");
        // window.location.replace("http://localhost:3000/producttable");
      } else {
        console.error("Upload failed.");
      }
    } catch (error) {
      console.error("Error:", error);
    }
  };

  const options = [
    { value: "Laptops", label: "Laptops" },
    { value: "smartphones", label: "smartphones" },
    { value: "Tablet", label: "Tablet" },
    { value: "Virtual Reality", label: "Virtual Reality" },
    // { value: "Basketball", label: "Basketball" },
    // { value: "Jordan", label: "Jordan" },
    // { value: "Golf", label: "Golf" },
    // { value: "Soccer", label: "Soccer" },
    // { value: "Tennis", label: "Tennis" },
    // { value: "Training", label: "Training" },
  ];

  return (
    <>
      <Navbar />
      <div className="container my-3 py-3">
        <h1 className="text-center">Create Product</h1>
        <hr />
        <div class="row my-4 h-100">
          <div className="col-md-4 col-lg-4 col-sm-8 mx-auto">
            <form>
              <div class="form my-3">
                <label for="Name">Title</label>
                <input
                  type="text"
                  class="form-control"
                  id="title"
                  placeholder="Product name"
                  value={title}
                  onChange={(e) => setTitle(e.target.value)}
                  required
                />
              </div>
              <div class="form my-3">
                <label for="text">Price</label>
                <input
                  type="text"
                  class="form-control"
                  id="price"
                  placeholder="Price"
                  value={price}
                  onChange={(e) => setPrice(e.target.value)}
                  required
                />
              </div>
              <div class="form my-3">
                <label for="category">Category</label>
                {/* <input
                  type="text"
                  class="form-control"
                  id="category"
                  placeholder="Category"
                /> */}
                <Select
                  options={options}
                  onChange={handleSelectChange}
                  value={options.find((option) => option.value === category)}
                  required
                />
              </div>

              <div class="form  my-3">
                <label for="description">Description</label>
                <textarea
                  rows={5}
                  class="form-control"
                  id="description"
                  placeholder="Enter description"
                  value={description}
                  onChange={(e) => setDescription(e.target.value)}
                  required
                />
              </div>

              <div class="form my-3">
                <label for="image">File Image</label>
                <input
                  type="text"
                  class="form-control"
                  id="image"
                  placeholder="Image link"
                  value={image}
                  onChange={(e) => setImage(e.target.value)}
                  required
                />
                <div style={{ marginTop: "3rem", marginBottom: "3rem" }}>
                  <center>
                    {selectedImage && (
                      <img
                        src={selectedImage}
                        alt="Selected"
                        style={{ maxWidth: "100%" }}
                      />
                    )}
                  </center>
                </div>
              </div>

              <div className="text-center">
                <button
                  class="my-2 px-4 mx-auto btn btn-primary"
                  type="submit"
                  onClick={handleSubmit}
                >
                  Add Product
                </button>
              </div>
            </form>
          </div>
        </div>
      </div>
      <Footer />
    </>
  );
};

export default CreateProduct;
