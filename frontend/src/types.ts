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
