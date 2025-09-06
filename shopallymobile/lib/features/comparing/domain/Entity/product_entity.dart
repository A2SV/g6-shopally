import 'package:equatable/equatable.dart';

import 'price_entity.dart';

class ProductEntity extends Equatable {
  final String id;
  final String title;
  final String imageUrl;
  final int aiMatchPercentage;
  final PriceEntity price;
  final double productRating;
  final String deliveryEstimate;
  final String description;
  final dynamic productSmallImageUrls;
  final int numberSold;
  final List<String> summaryBullets;
  final String deepLinkUrl;
  final double tax;
  final double discount;


  const ProductEntity({
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

  @override
  List<Object?> get props => [
    id,
    title,
    imageUrl,
    aiMatchPercentage,
    price,
    productRating,
    deliveryEstimate,
    description,
    productSmallImageUrls,
    numberSold,
    summaryBullets,
    deepLinkUrl,
    tax,
    discount,
  ];
}
