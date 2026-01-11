export interface Article {
  id: string;
  //
  title: string;
  description?: string;
  content?: string;
  //
  author: string;
  published: boolean;
  //
  created_at: string;
  modifited_at: string;
}

export type Comment = {
  id: number;
  //
  article_id: number;
  author_name: string;
  content: string;
  //
  created_at: string;
  modifited_at: string;
};
