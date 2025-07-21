import { Link } from "react-router";

export const NavMenu = () => {
  return (
    <div className="navbar bg-base-100 shadow-sm">
      <div className="flex-1">
        <Link to={"/"} className="btn btn-neutral text-xl">
          Heartbit
        </Link>
      </div>
      <div className="flex-none">
        <ul className="menu menu-horizontal px-1">
          <li>
            <Link to={"/article/write"} className="btn btn-success">
              New
            </Link>
          </li>
          {/*
            <li>
                <details>
                <summary>Parent</summary>
                <ul className="bg-base-100 rounded-t-none p-2">
                    <li>
                    <a>Link 1</a>
                    </li>
                    <li>
                    <a>Link 2</a>
                    </li>
                </ul>
                </details>
            </li>
          */}
        </ul>
      </div>
    </div>
  );
};
