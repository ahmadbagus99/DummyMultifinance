import { useState } from "react";
import { useNavigate } from "react-router-dom";
import logo from "../images/DummyMultifinance.png";
import slider1 from "../images/slider/1.jpg";

const Login = () => {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [error, setError] = useState("");
    const navigate = useNavigate();

    const handleLogin = async (e) => {
        e.preventDefault();
        setError("");
    
        try {
            const response = await fetch("http://localhost:8096/login", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({ username, password }),
                mode: "cors",
            });
        
            if (response.status === 204) {
                throw new Error("Server tidak mengembalikan response.");
            }
            if (!response.ok) {
                throw new Error(`HTTP error! Status: ${response.status}`);
            }
    
            const result = await response.json();
            if (result.resultCode === "00") {
                localStorage.setItem("token", result.data.accessToken);
                localStorage.setItem("expired", result.data.exp);
                navigate("/dashboard", { replace: true });
            } else {
                setError(result.message || "Login gagal, coba lagi!");
            }
        } catch (err) {
            console.error("Login error:", err);
            setError("Terjadi kesalahan: " + err.message);
        }
    };
    

    return (
        <section className="flex h-screen">
            {/* Bagian Kiri (Form Login) */}
            <div className="w-1/2 flex flex-col justify-center items-center bg-gray-900 text-white">
                <img src={logo} alt="Logo Asta Karya" className="w-64 mb-6" />
                <form className="w-1/3" onSubmit={handleLogin}>
                    <h3 className="text-center text-lg font-semibold mb-4">Log in</h3>
                    {error && <p className="text-red-500 text-center">{error}</p>}
                    <input
                        type="text"
                        name="username"
                        className="w-full p-2 mb-4 border rounded bg-gray-800 text-white"
                        placeholder="Username"
                        value={username}
                        onChange={(e) => setUsername(e.target.value)}
                        required
                    />
                    <input
                        type="password"
                        name="password"
                        className="w-full p-2 mb-4 border rounded bg-gray-800 text-white"
                        placeholder="Password"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                        required
                    />
                    <button className="w-full bg-blue-500 text-white py-2 rounded">Login</button>
                </form>
            </div>

            {/* Bagian Kanan (Gambar) */}
            <div className="w-1/2 hidden sm:block">
                <img src={slider1} alt="Login Background" className="w-full h-full object-cover" />
            </div>
        </section>
    );
};

export default Login;
