import { Navigate } from "react-router-dom";

const ProtectedRoute = ({ children }) => {
    const token = localStorage.getItem("token");
    const expiredStr = localStorage.getItem("expired");

    // Ambil waktu sekarang dalam zona WIB
    const date = new Date();
    const options = { timeZone: "Asia/Jakarta" };
    const dateInWIB = new Date(date.toLocaleString("en-US", options));

    const year = dateInWIB.getFullYear();
    const month = String(dateInWIB.getMonth() + 1).padStart(2, "0");
    const day = String(dateInWIB.getDate()).padStart(2, "0");
    const hours = String(dateInWIB.getHours()).padStart(2, "0");
    const minutes = String(dateInWIB.getMinutes()).padStart(2, "0");
    const seconds = String(dateInWIB.getSeconds()).padStart(2, "0");

    const formattedDate = `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`;

    function parseDateTime(dateTimeStr) {
        if (!dateTimeStr) return null;
        const [datePart, timePart] = dateTimeStr.split(" ");
        const [year, month, day] = datePart.split("-").map(Number);
        const [hour, minute, second] = timePart.split(":").map(Number);
        return new Date(year, month - 1, day, hour, minute, second);
    }

    const expiredDate = parseDateTime(expiredStr);

    if (!token || (expiredDate && dateInWIB >= expiredDate)) {
        console.log("Token expired, menghapus dari localStorage...");
        localStorage.removeItem("token");
        localStorage.removeItem("expired");

        // Redirect ke login
        return <Navigate to="/login" replace />;
    }

    return children;
};

export default ProtectedRoute;
