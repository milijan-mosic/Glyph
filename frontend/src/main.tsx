import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import "./index.css";
import { ListArticles } from "@pages/article/List.tsx";
import { ReadArticle } from "@pages/article/Read.tsx";
import { WriteOrEditArticle } from "@pages/article/WriteOrEdit.tsx";
import { BrowserRouter, Routes, Route } from "react-router";
import { NavMenu } from "@components/menu/NavMenu";
import flowers from "@assets/flowers.jpg";

const root: HTMLElement | null = document.getElementById("root");

createRoot(root!).render(
  <StrictMode>
    <BrowserRouter>
      <img src={flowers} className="w-full h-fit -z-10 absolute" />
      <NavMenu />

      <Routes>
        <Route path="/" element={<ListArticles />} />
        <Route path="/article/:articleId" element={<ReadArticle />} />
        <Route path="/article/write" element={<WriteOrEditArticle />} />
        <Route
          path="/article/edit/:articleId"
          element={<WriteOrEditArticle />}
        />
      </Routes>
    </BrowserRouter>
  </StrictMode>
);
