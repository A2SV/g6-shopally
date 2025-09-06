import 'package:equatable/equatable.dart';

class Product extends Equatable {
  final String id;
  final String title;
  final String imageUrl;
  final double price;
  final int minOrder;
  final double rating;
  final int iscompare;
  final int issaved;

  const Product({
    required this.id,
    required this.title,
    required this.imageUrl,
    required this.price,
    required this.minOrder,
    required this.rating,
    required this.iscompare,
    required this.issaved,
  });

  @override
  List<Object> get props => [
    id,
    title,
    imageUrl,
    price,
    minOrder,
    rating,
    iscompare,
    issaved,
  ];
}
