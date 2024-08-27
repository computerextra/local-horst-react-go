import { useEffect } from "react";
import { useLocation } from "react-router-dom";

function active(title: string, should: string): boolean {
  if (title.includes(should)) {
    return true;
  } else {
    return false;
  }
}

export default function Menu() {
  const location = useLocation();

  useEffect(() => {
    const elem = document.querySelector("#Start");
    if (elem == null) return;
    console.log(location.pathname);
    if (!active(location.pathname, "/")) {
      elem.classList.remove("border");
      elem.classList.remove("border-primary-light");
    }
  }, []);

  return (
    <div className="h-screen w-60 rounded-br-md rounded-tr-lg bg-slate-800 p-4 antialiased print:hidden">
      <div className="flex items-center gap-2">
        <div>
          <p className="font-medium text-white">Local Horst</p>
        </div>
      </div>
      <ul className="mt-10 flex w-full flex-col gap-3">
        <li>
          <a
            href="/"
            id="Start"
            className={`flex items-center gap-2 rounded-md bg-primary-light/40 px-3 py-2 text-primary-white border border-primary-light"`}
          >
            <svg
              className="h-5 w-5 fill-white"
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 24 24"
            >
              <path d="M5 3V19H21V21H3V3H5ZM20.2929 6.29289L21.7071 7.70711L16 13.4142L13 10.415L8.70711 14.7071L7.29289 13.2929L13 7.58579L16 10.585L20.2929 6.29289Z"></path>
            </svg>
            Startseite
          </a>
        </li>
      </ul>
    </div>
  );
}
