import { useEffect, useState } from "react";
// import logoDark from "../images/asta-karya-white.png";
// import logoLight from "../images/asta-karya.png";
import logoDark from "../images/newLight.png";
import logoLight from "../images/newLight.png";

const Logo = () => {
  const [isDarkMode, setIsDarkMode] = useState(false);

  useEffect(() => {
    const checkDarkMode = () => setIsDarkMode(document.documentElement.classList.contains("dark"));
    checkDarkMode();
    const observer = new MutationObserver(checkDarkMode);
    observer.observe(document.documentElement, { attributes: true, attributeFilter: ["class"] });

    return () => observer.disconnect();
  }, []);

  return <img src={isDarkMode ? logoDark : logoLight} alt="Logo" width={200} height={32} />;
};

export default Logo;
