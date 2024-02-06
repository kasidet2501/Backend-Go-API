// import React, { useState, useEffect } from "react";
// import Cookies from "js-cookie";
// import axios from "axios";
// import { Footer, Navbar } from "../components";
// import Skeleton from "react-loading-skeleton";
// import "react-loading-skeleton/dist/skeleton.css";

// const Cart = () => {
//   const [cart, setCart] = useState([]);
//   const [loading, setLoading] = useState(false);
//   let componentMounted = true;

//   // useEffect(() => {
//   //   const headers = {
//   //     "Content-Type": "application/json",
//   //     authorization: `Bearer ${Cookies.get("jwt")}`,
//   //   };

//   //   axios
//   //     .get("https://127.0.0.1:8080/cart/", { headers })
//   //     .then((response) => {
//   //       console.log(response);
//   //       setCart(response);
//   //     })
//   //     .catch((err) => {
//   //       console.log(err);
//   //     });
//   //   console.log("cart : ", cart);
//   // }, [setCart]);

//   useEffect(() => {
//     const getCarts = async () => {
//       setLoading(true);
//       const response = await fetch("http://127.0.0.1:8080/cart", {
//         method: "GET",
//         headers: {
//           "Content-Type": "application/json",
//           authorization: `Bearer ${Cookies.get("jwt")}`,
//           // Authorization : `Bearer ${Cookies.get('jwt')}`
//         },
//         // credentials: "include",
//       });
//       if (componentMounted) {
//         const data = await response.json();
//         setCart(data);
//         setLoading(false);
//       }
//       return () => {
//         componentMounted = false;
//       };
//     };

//     getCarts();
//   }, []);

//   const Loading = () => {
//     return (
//       <>
//         <div className="col-12 py-5 text-center">
//           <Skeleton height={40} width={560} />
//         </div>
//         <div className="col-md-4 col-sm-6 col-xs-8 col-12 mb-4">
//           <Skeleton height={592} />
//         </div>
//         <div className="col-md-4 col-sm-6 col-xs-8 col-12 mb-4">
//           <Skeleton height={592} />
//         </div>
//         <div className="col-md-4 col-sm-6 col-xs-8 col-12 mb-4">
//           <Skeleton height={592} />
//         </div>
//         <div className="col-md-4 col-sm-6 col-xs-8 col-12 mb-4">
//           <Skeleton height={592} />
//         </div>
//         <div className="col-md-4 col-sm-6 col-xs-8 col-12 mb-4">
//           <Skeleton height={592} />
//         </div>
//         <div className="col-md-4 col-sm-6 col-xs-8 col-12 mb-4">
//           <Skeleton height={592} />
//         </div>
//       </>
//     );
//   };

//   const ShowProduct = () => {
//     {
//       cart.map((item) => {
//         // let img = product.image;
//         return (
//           <div
//             id={item._id}
//             key={item._id}
//             className="col-md-4 col-sm-6 col-xs-8 col-12 mb-4"
//           >
//             <div className="card text-center h-100" key={item._id}>
//               <img
//                 className="card-img-top p-3"
//                 // src={require("../images/" + img)}
//                 alt="Card"
//                 height={300}
//               />
//               <div className="card-body">
//                 <h5 className="card-title">{item.productId}...</h5>
//                 <p className="card-text">{item.quantity}...</p>
//               </div>

//               <ul className="list-group list-group-flush">
//                 <li className="list-group-item lead">$ {item.price}</li>
//                 {/* <li className="list-group-item">Dapibus ac facilisis in</li>
//                     <li className="list-group-item">Vestibulum at eros</li> */}
//               </ul>
//               {/* <div className="card-body">
//                   <Link
//                     to={"/product/" + product.id}
//                     className="btn btn-dark m-1"
//                   >
//                     Buy Now
//                   </Link>
//                   <button
//                     className="btn btn-dark m-1"
//                     onClick={() => addProduct(product)}
//                   >
//                     Add to Cart
//                   </button>
//                 </div> */}
//             </div>
//           </div>
//         );
//       });
//     }
//   };

//   return (
//     <>
//       <Navbar />
//       <div className="container my-3 py-3">
//         <h1 className="text-center">Cart</h1>
//         <hr />
//         {loading ? <Loading /> : <ShowProduct />}
//       </div>
//       <Footer />
//     </>
//   );
// };

// export default Cart;
