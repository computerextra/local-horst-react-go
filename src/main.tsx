import React from "react";
import ReactDOM from "react-dom/client";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import "./index.css";
import Home from "./Pages/Home";
import Menu from "./Layout/Menu";

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <BrowserRouter>
      <div className="flex">
        <Menu />
        <div className="flex-1 px-4 py-6 hidden sm:block">
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/Test" element={<Home />} />
          </Routes>
        </div>
      </div>
    </BrowserRouter>
  </React.StrictMode>,
);
