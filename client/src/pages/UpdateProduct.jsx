import React, { useEffect, useState } from "react";
import Cookies from "js-cookie";
import axios from "axios";
import { Footer, Navbar, Product } from "../components";
import Select from "react-select";
import Skeleton from "react-loading-skeleton";
import { useParams } from "react-router-dom";

const UpdateProduct = () => {
  const { id } = useParams();

  const [selectedImage, setSelectedImage] = useState(null);
  const [title, setTitle] = useState("");
  const [price, setPrice] = useState("");
  const [description, setDescription] = useState("");
  const [category, setCategory] = useState("");
  const [image, setImage] = useState("");

  let componentMounted = true;

  useEffect(() => {
    const getProduct = async () => {
      const response = await fetch(`http://127.0.0.1:8080/user/product/${id}`, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          authorization: `Bearer ${Cookies.get("jwt")}`,
        },
      });
      if (componentMounted) {
        const data = await response.json();
        setTitle(data.title);
        setPrice(data.price);
        setDescription(data.description);
        setCategory(data.category);
        setImage(data.image);
      }
      return () => {
        componentMounted = false;
      };
    };
    getProduct();
  }, [id]);

  const handleSelectChange = (selectedOption) => {
    setCategory(selectedOption.value);
  };

  // async function handleSubmit() {
  const handleSubmit = async () => {
    try {
      const priceParse = parseFloat(price);

      const values = {
        title: title,
        price: priceParse,
        description: description,
        category: category,
        image: image,
      };

      let response = await fetch(`http://127.0.0.1:8080/product/${id}`, {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
          authorization: `Bearer ${Cookies.get("jwt")}`,
        },
        body: JSON.stringify(values),
      });
      window.location.replace("http://localhost:3000/producttable");
      if (response.status === 200) {
        console.log("Upload successful!");
        window.location.replace("http://localhost:3000/producttable");
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
      {/* {!loading ? <ShowForm /> : <ShowForm />} */}
      <div className="container my-3 py-3">
        <h1 className="text-center">Update Product</h1>
        <hr />
        <div class="row my-4 h-100">
          <div className="col-md-4 col-lg-4 col-sm-8 mx-auto">
            <form onSubmit={handleSubmit}>
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
                  class="my-2 px-4 mx-auto btn btn-dark"
                  type="submit"
                  // onClick={handleSubmit}
                >
                  Update
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

export default UpdateProduct;
