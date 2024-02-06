import React, { useState, useEffect } from "react";
import Cookies from "js-cookie";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { Footer, Navbar } from "../components";
import { Link } from "react-router-dom";

const ProductTable = () => {
  const [data, setData] = useState([]);

  let componentMounted = true;

  useEffect(() => {
    const getProducts = async () => {
      // setLoading(true);
      const response = await fetch("http://127.0.0.1:8080/user/product", {
        headers: {
          "Content-Type": "application/json",
          authorization: `Bearer ${Cookies.get("jwt")}`,
          // Authorization : `Bearer ${Cookies.get('jwt')}`
        },
        // credentials: "include",
      });
      if (componentMounted) {
        setData(await response.clone().json());

        // setLoading(false);
      }

      return () => {
        componentMounted = false;
      };
    };

    getProducts();
  }, [handleDelete]);

  async function handleDelete(id) {
    try {
      let response = await fetch(`http://127.0.0.1:8080/product/${id}`, {
        method: "DELETE",
        headers: {
          "Content-Type": "application/json",
          authorization: `Bearer ${Cookies.get("jwt")}`,
        },
      });

      if (response.status === 200) {
        console.log("Delete successful!");
      } else {
        console.error("Delete failed.");
      }
    } catch (error) {
      console.error("Error:", error);
    }
  }

  return (
    <>
      <Navbar />

      <div className="container mt-5">
        <div className="col-12" style={{ marginBottom: "3rem" }}>
          <h2 className="display-5 text-center">Products</h2>
          <hr />
        </div>
        <div className="text-start" style={{ marginBottom: "1rem" }}>
          <Link to={"/createproduct/"} className="btn btn-primary">
            + Add Product
          </Link>
        </div>
        <table className="table">
          <thead>
            <tr>
              <th scope="col1">
                <center>IMAGE</center>
              </th>
              <th scope="col2">Title</th>
              <th scope="col3">Price</th>
              <th scope="col4">Category</th>
              <th scope="col5">
                <center>Action</center>
              </th>
            </tr>
          </thead>
          <tbody>
            {data.map((item) => {
              let img = item.image;
              return (
                <tr key={item.id}>
                  <td>
                    <center>
                      <img
                        // className="card-img-top p-3"
                        // src={require("../images/" + img)}
                        src={img}
                        // alt="Card"
                        height={50}
                      />
                    </center>
                  </td>
                  <td>{item.title}</td>
                  <td>{item.price}</td>
                  <td>{item.category}</td>
                  <td>
                    <center>
                      <Link
                        className="btn btn-success"
                        to={"/updateproduct/" + item.id}
                        // onClick={() => addtoCart(product.id)}
                      >
                        UPDATE
                      </Link>
                      <Link
                        className="btn btn-danger"
                        style={{ marginLeft: "1rem" }}
                        onClick={() => handleDelete(item.id)}
                      >
                        DELETE
                      </Link>
                    </center>
                  </td>
                </tr>
              );
            })}
          </tbody>
        </table>
      </div>

      <Footer />
    </>
  );
};

export default ProductTable;
