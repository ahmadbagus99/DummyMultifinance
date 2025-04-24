import React from "react";
import { Link, useNavigate } from "react-router-dom";
import Swal from "sweetalert2";

const LogoutButton = ({ setDropdownOpen }) => {
  const navigate = useNavigate();

  const handleLogout = (e) => {
    e.preventDefault(); // Mencegah navigasi langsung ke /login
  
    Swal.fire({
      title: "Apakah Anda yakin ingin logout?",
      text: "Anda akan keluar dari sesi ini!",
      icon: "warning",
      showCancelButton: true,
      confirmButtonText: "Ya, Logout",
      cancelButtonText: "Batal",
    }).then((result) => {
      if (result.isConfirmed) {
        // Hapus token dari localStorage atau sessionStorage
        localStorage.removeItem("token"); 
        sessionStorage.removeItem("token");
  
        // Atau jika menggunakan cookies, hapus token dengan mengatur waktu kedaluwarsa ke masa lalu
        document.cookie = "token=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
  
        console.log("User logged out, token removed");
        setDropdownOpen(false);
        navigate("/login"); // Arahkan ke halaman login setelah logout
      }
    });
  };
  

  return (
    <li>
      <Link
        className="font-medium text-sm text-violet-500 hover:text-violet-600 dark:hover:text-violet-400 flex items-center py-1 px-3"
        to="/login"
        onClick={handleLogout} // Ganti fungsi ini
      >
        Log Out
      </Link>
    </li>
  );
};

export default LogoutButton;
