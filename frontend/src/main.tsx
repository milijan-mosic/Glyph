import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import "./index.css";
import List from "./pages/article/List.tsx";
import Read from "./pages/article/Read.tsx";
import WriteOrEdit from "./pages/article/WriteOrEdit.tsx";
import { BrowserRouter, Routes, Route } from "react-router";

const root: HTMLElement | null = document.getElementById("root");

createRoot(root!).render(
  <StrictMode>
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<List />} />
        <Route path="/article/:articleId" element={<Read />} />
        <Route path="/article/write" element={<WriteOrEdit />} />
        <Route path="/article/edit/:articleId" element={<WriteOrEdit />} />
      </Routes>
    </BrowserRouter>
  </StrictMode>
);
