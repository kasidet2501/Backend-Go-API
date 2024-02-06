import { useEffect, useState } from "react";
import Cookies from "js-cookie";
import { Navbar, Main, Product, Footer } from "../components";
import { Register } from "../pages";

function Home() {
  // const [role, setRole] = useState("");
  // useEffect(() => {
  //   const getRole = async () => {
  //     const response = await fetch("http://127.0.0.1:8080/role", {
  //       headers: {
  //         "Content-Type": "application/json",
  //         authorization: `Bearer ${Cookies.get("jwt")}`,
  //         // Authorization : `Bearer ${Cookies.get('jwt')}`
  //       },
  //       // credentials: "include",
  //     });
  //     const data = await response.json();
  //     console.log("Data : ", data);
  //     setRole(data);
  //     console.log("Role : ", role);
  //   };
  //   getRole();
  // }, [role]);

  return (
    <>
      <Navbar />
      <Main />
      <Product />
      <Footer />
    </>
  );
}

export default Home;
