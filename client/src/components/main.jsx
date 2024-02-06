import React from "react";

const Home = () => {
  return (
    <>
      <div className="hero border-1 pb-3">
        <div className="card bg-dark text-white border-0 mx-3">
          <img
            className="card-img img-fluid"
            src="./assets/main.png.jpg"
            // src="https://rtlimages.apple.com/cmc/dieter/store/16_9/R733.png?resize=672:378&output-format=jpg&output-quality=85&interpolation=progressive-bicubic"
            alt="Card"
            height={500}
          />
          <div className="card-img-overlay d-flex align-items-center">
            <div className="container">
              <h5 className="card-title fs-1 text fw-lighter">Go Ecommerce</h5>
              <p className="card-text fs-5 d-none d-sm-block ">
                Go beyond the ordinary in e-commerce. Experience the
                extraordinary with enhanced features for a seamless online
                shopping journey.
              </p>
            </div>
          </div>
        </div>
      </div>
    </>
  );
};

export default Home;
