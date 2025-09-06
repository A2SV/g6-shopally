import 'package:shopallymobile/features/saveditem/domain/entities/product.dart';

class ProductModel extends Product {
  const ProductModel({
    // These should be required to match the Product entity
    required super.id,
    required super.title,
    required super.imageUrl,
    required super.price,
    required super.minOrder,
    required super.rating,
    int? iscompare,
    int? issaved,
    // This super call is what was missing. It passes the values to the parent Product class.
  }) : super(iscompare: iscompare ?? 0, issaved: issaved ?? 0);

  factory ProductModel.fromEntity(Product product) {
    return ProductModel(
      id: product.id,
      title: product.title,
      imageUrl: product.imageUrl,
      price: product.price,
      minOrder: product.minOrder,
      rating: product.rating,
      iscompare: product.iscompare,
      issaved: product.issaved,
    );
  }

  factory ProductModel.fromMap(Map<String, dynamic> map) {
    return ProductModel(
      id: map['id'] as String? ?? '',
      title: map['title'] as String? ?? '',
      imageUrl: map['imageUrl'] as String? ?? '',
      price: (map['price'] as num?)?.toDouble() ?? 0.0,
      minOrder: int.tryParse(map['minOrder']?.toString() ?? '0') ?? 0,
      rating: (map['rating'] as num?)?.toDouble() ?? 0.0,
      iscompare: int.tryParse(map['iscompare']?.toString() ?? '0') ?? 0,
      issaved: int.tryParse(map['issaved']?.toString() ?? '0') ?? 0,
    );
  }

  // This method is needed to insert data into the database.
  Map<String, dynamic> toMap() {
    return {
      'id': id,
      'title': title,
      'imageUrl': imageUrl,
      'price': price,
      'minOrder': minOrder,
      'rating': rating,
      'iscompare': iscompare,
      'issaved': issaved,
    };
  }

  @override
  String toString() {
    return 'ProductModel(id: $id, title: $title, price: $price, minOrder: $minOrder, rating: $rating, imageUrl: $imageUrl)';
  }
}
