import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import "./index.css";
import { ListArticles } from "@pages/article/List.tsx";
import { ReadArticle } from "@pages/article/Read.tsx";
import WriteOrEdit from "@pages/article/WriteOrEdit.tsx";
import { BrowserRouter, Routes, Route } from "react-router";
import { NavMenu } from "@components/menu/NavMenu";

const root: HTMLElement | null = document.getElementById("root");

createRoot(root!).render(
  <StrictMode>
    <BrowserRouter>
      <NavMenu />

      <Routes>
        <Route path="/" element={<ListArticles />} />
        <Route path="/article/:articleId" element={<ReadArticle />} />
        <Route path="/article/write" element={<WriteOrEdit />} />
        <Route path="/article/edit/:articleId" element={<WriteOrEdit />} />
      </Routes>
    </BrowserRouter>
  </StrictMode>
);
