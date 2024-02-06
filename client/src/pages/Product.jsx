import React, { useEffect, useState } from "react";
import Skeleton from "react-loading-skeleton";
import Cookies from "js-cookie";
import { Link, useParams } from "react-router-dom";
import Marquee from "react-fast-marquee";
import { useDispatch } from "react-redux";
import { addCart } from "../redux/action";

import { Footer, Navbar } from "../components";

const Product = () => {
  const { id } = useParams();
  const [product, setProduct] = useState([]);
  const [loading, setLoading] = useState(false);
  const [role, setRole] = useState("");

  const dispatch = useDispatch();

  const addProduct = (product) => {
    dispatch(addCart(product));
  };

  useEffect(() => {
    const getProduct = async () => {
      setLoading(true);
      // const response = await fetch("http://127.0.0.1:8080/user/product", {
      //   credentials: "include",
      // });
      const response = await fetch(`http://127.0.0.1:8080/user/product/${id}`, {
        headers: {
          "Content-Type": "application/json",
          authorization: `Bearer ${Cookies.get("jwt")}`,
          // Authorization : `Bearer ${Cookies.get('jwt')}`
        },
      });
      const data = await response.json();
      setProduct(data);
      setLoading(false);
    };

    const getRole = async () => {
      const response = await fetch("http://127.0.0.1:8080/role", {
        headers: {
          "Content-Type": "application/json",
          authorization: `Bearer ${Cookies.get("jwt")}`,
          // Authorization : `Bearer ${Cookies.get('jwt')}`
        },
        // credentials: "include",
      });
      const data = await response.json();
      setRole(data);
    };
    getProduct();
    getRole();
  }, [id]);

  async function addtoCart(id) {
    try {
      await fetch("http://127.0.0.1:8080/cart", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          authorization: `Bearer ${Cookies.get("jwt")}`,
        },
        credentials: "include",
        body: JSON.stringify({
          productId: id,
        }),
      });
    } catch (error) {
      console.error("Error add item to cart failed:", error);
    }
  }

  const Loading = () => {
    return (
      <>
        <div className="container my-5 py-2">
          <div className="row">
            <div className="col-md-6 py-3">
              <Skeleton height={400} width={400} />
            </div>
            <div className="col-md-6 py-5">
              <Skeleton height={30} width={250} />
              <Skeleton height={90} />
              <Skeleton height={40} width={70} />
              <Skeleton height={50} width={110} />
              <Skeleton height={120} />
              <Skeleton height={40} width={110} inline={true} />
              <Skeleton className="mx-3" height={40} width={110} />
            </div>
          </div>
        </div>
      </>
    );
  };

  const ShowProduct = () => {
    let img = product.image;
    return (
      <>
        <div className="container my-5 py-2">
          <div className="row">
            <div className="col-md-6 col-sm-12 py-3">
              <img
                className="img-fluid"
                // src={require("../images/" + img)}
                src={img}
                alt={product.title}
                width="400px"
                height="400px"
              />
            </div>
            <div className="col-md-6 col-md-6 py-5">
              <h4 className="text-uppercase text-muted">{product.category}</h4>
              <h1 className="display-5">{product.title}</h1>
              <h3 className="display-6  my-4">${product.price}</h3>
              <p className="lead">{product.description}</p>

              {role == "admin" ? "" : ""}

              {role == "admin" ? "" : ""}

              {role == "user" ? (
                <button
                  className="btn btn-outline-dark"
                  // onClick={() => addtoCart(product.id)}
                  onClick={() => addProduct(product)}
                >
                  Add to Cart
                </button>
              ) : (
                ""
              )}

              {role == "user" ? (
                <Link to="/cart" className="btn btn-dark mx-3">
                  Go to Cart
                </Link>
              ) : (
                ""
              )}
            </div>
          </div>
        </div>
      </>
    );
  };

  return (
    <>
      <Navbar />
      <div className="container">
        <div className="row">
          {!product.image ? <Loading /> : <ShowProduct />}
        </div>
      </div>

      <Footer />
    </>
  );
};

export default Product;
