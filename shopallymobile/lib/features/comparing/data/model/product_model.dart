import 'dart:convert';


import 'package:shopallymobile/features/comparing/data/model/price_model.dart';
import 'package:shopallymobile/features/comparing/domain/Entity/product_entity.dart';

class ProductModel {
  final String id;
  final String title;
  final String imageUrl;
  final int aiMatchPercentage;
  final PriceModel price;
  final double productRating;
  final String deliveryEstimate;
  final String description;
  final dynamic productSmallImageUrls;
  final int numberSold;
  final List<String> summaryBullets;
  final String deepLinkUrl;
  final double tax;
  final double discount;

  const ProductModel({
    required this.imageUrl,
    required this.price,
    required this.id,
    required this.title,
    required this.aiMatchPercentage,
    required this.productRating,
    required this.deepLinkUrl,
    required this.deliveryEstimate,
    required this.description,
    required this.productSmallImageUrls,
    required this.numberSold,
    required this.summaryBullets,
    required this.tax,
    required this.discount,
  });

  factory ProductModel.fromJson(Map<String, dynamic> json) {
    return ProductModel(
      imageUrl: json['imageUrl'],
      price: PriceModel.fromJson(json['price']),
      summaryBullets: List<String>.from(json['summaryBullets']),
      id: json['id'],
      title: json['title'],
      aiMatchPercentage: (json['aiMatchPercentage'] as num).toInt(),
      productRating: (json['productRating'] as num).toDouble(),
      deepLinkUrl: json['deeplinkUrl'],
      deliveryEstimate: json['deliveryEstimate'],
      description: json['description'],
      numberSold: (json['numberSold'] as num).toInt(),
      tax: (json['taxRate'] as num).toDouble(),
      discount: (json['discount'] as num).toDouble(),
      productSmallImageUrls: json['productSmallImageUrls'],
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'imageUrl': imageUrl,
      'price': price.toJson(),
      'id': id,
      'title': title,
      'aiMatchPercentage': aiMatchPercentage,
      'productRating': productRating,
      'deeplinkUrl': deepLinkUrl,
      'deliveryEstimate': deliveryEstimate,
      'description': description,
      'productSmallImageUrls': productSmallImageUrls,
      'numberSold': numberSold,
      'summaryBullets': summaryBullets,
      'taxRate': tax,
      'discount': discount,
    };
  }

  ProductEntity toEntity() {
    return ProductEntity(
      imageUrl: imageUrl,
      price: price.toEntity(price),
      id: id,
      title: title,
      aiMatchPercentage: aiMatchPercentage,
      productRating: productRating,
      deepLinkUrl: deepLinkUrl,
      deliveryEstimate: deliveryEstimate,
      description: description,
      productSmallImageUrls: productSmallImageUrls,
      numberSold: numberSold,
      summaryBullets: summaryBullets,
      tax: tax,
      discount: discount,
    );
  }

  /// ✅ Safe conversion for numeric types
  factory ProductModel.fromEntity(ProductEntity product) {
    return ProductModel(
      imageUrl: product.imageUrl,
      price: PriceModel.fromEntity(product.price),
      id: product.id,
      title: product.title,
      aiMatchPercentage: (product.aiMatchPercentage as num).toInt(),
      productRating: (product.productRating as num).toDouble(),
      deepLinkUrl: product.deepLinkUrl,
      deliveryEstimate: product.deliveryEstimate,
      description: product.description,
      productSmallImageUrls: product.productSmallImageUrls,
      numberSold: (product.numberSold as num).toInt(),
      summaryBullets: product.summaryBullets,
      tax: (product.tax as num).toDouble(),
      discount: (product.discount as num).toDouble(),
    );
  }

  /// ✅ Safe parsing when restoring from DB (stored as strings/JSON)
  factory ProductModel.fromDb(Map<String, dynamic> json) {
    return ProductModel(
      imageUrl: json['imageUrl'],
      price: PriceModel.fromJson(jsonDecode(json['price'])),
      id: json['id'],
      title: json['title'],
      aiMatchPercentage: (json['aiMatchPercentage'] as num).toInt(),
      productRating: (json['productRating'] as num).toDouble(),
      deepLinkUrl: json['deepLinkUrl'],
      deliveryEstimate: json['deliveryEstimate'],
      description: json['description'],
      productSmallImageUrls: jsonDecode(json['productSmallImageUrls']),
      numberSold: (json['numberSold'] as num).toInt(),
      summaryBullets: List<String>.from(jsonDecode(json['summaryBullets'])),
      tax: (json['tax'] as num).toDouble(),
      discount: (json['discount'] as num).toDouble(),
    );
  }
}
