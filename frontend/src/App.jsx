import React, { useEffect } from "react";
import { Routes, Route, useLocation, Navigate } from "react-router-dom";
import ProtectedRoute from "./config/ProtectedRoute";
// import "bootstrap/dist/css/bootstrap.min.css";
import "./css/style.css";
import "./charts/ChartjsConfig";

// Import pages
import Layout from "./layout/Layout"; // Tambahkan layout utama
import Dashboard from "./pages/Dashboard";
import MasterDataSiswa from "./pages/MasterDataSiswa";
// import Akademik from "./pages/Akademik";
// import DokumenKontrak from "./pages/DokumenKontrak";
import LoginPage from "./pages/Login";

function App() {
  const location = useLocation();

  useEffect(() => {
    document.querySelector("html").style.scrollBehavior = "auto";
    window.scroll({ top: 0 });
    document.querySelector("html").style.scrollBehavior = "";
  }, [location.pathname]); // Triggered on route change

  return (
    <Routes>
      {/* Rute Login (tidak butuh proteksi) */}
      <Route path="/login" element={<LoginPage />} />

      {/* Semua halaman dashboard berada di dalam ProtectedRoute */}
      <Route
        path="/"
        element={
          <ProtectedRoute>
            <Layout />
          </ProtectedRoute>
        }
      >
        {/* Default route jika user mengunjungi "/" */}
        <Route index element={<Dashboard />} />
        <Route path="dashboard" element={<Dashboard />} />
        <Route path="master-data-siswa" element={<MasterDataSiswa />} />
        {/* <Route path="akademik" element={<Akademik />} />
        <Route path="dokumen-kontrak" element={<DokumenKontrak />} /> */}
      </Route>

      {/* Redirect jika rute tidak ditemukan */}
      <Route path="*" element={<Navigate to="/dashboard" />} />
    </Routes>
  );
}

export default App;
