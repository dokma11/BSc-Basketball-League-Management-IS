export interface WishlistAsset {
  idZeljTim: number;
	datDodZeljTim: Date; // Date of creation
	belesZeljTim: string;  // Notes
	idTipZelje: number; // Wishlist Asset Type foreign key
	idPrava: number; // Draft Rights foreign key
	idPik: number; // Pick foreign key
	idIgrac: number; // Player foreign key
	idTim: number; // Team foreign key
}
