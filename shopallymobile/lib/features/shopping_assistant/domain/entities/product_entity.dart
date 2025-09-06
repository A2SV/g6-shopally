import 'package:equatable/equatable.dart';

class ProductEntity extends Equatable {
  final String id;
  final String productName;
  final String productDescription;
  final String imageUrl;

  // Pricing (multi–currency)
  final double priceUsd;
  final double priceEtb;
  final DateTime? fxTimestamp;

  // Scores / ratings
  final double aiMatchPercentage;
  final double productRating; // 0–100?
  final double sellerScore;
  final double taxRate;
  final double discount; // percent or absolute depending on source

  // Inventory / sales
  final bool inStock;
  final int numberSold;

  // Content
  final String deliveryEstimate;
  final String customerHighlights;
  final String customerReview;
  final List<String> summaryBullets;
  final String deeplinkUrl;

  const ProductEntity({
    required this.id,
    required this.productName,
    required this.productDescription,
    required this.imageUrl,
    required this.priceUsd,
    this.priceEtb = 0,
    this.fxTimestamp,
    this.aiMatchPercentage = 0,
    this.productRating = 0,
    this.sellerScore = 0,
    this.taxRate = 0,
    this.discount = 0,
    this.inStock = true,
    this.numberSold = 0,
    this.deliveryEstimate = '',
    this.customerHighlights = '',
    this.customerReview = '',
    this.summaryBullets = const [],
    this.deeplinkUrl = '',
  });

  factory ProductEntity.fromMap(Map<String, dynamic> map) {
    final price = map['price'] as Map<String, dynamic>?;

    DateTime? ts;
    final rawTs = price?['fxTimestamp'];
    if (rawTs is String) {
      try {
        ts = DateTime.parse(rawTs);
      } catch (_) {}
    }

    return ProductEntity(
      id: map['id']?.toString() ?? '',
      productName: map['title']?.toString() ??
          map['productName']?.toString() ??
          map['name']?.toString() ??
          '',
      productDescription: map['description']?.toString() ??
          map['productDescription']?.toString() ??
          '',
      imageUrl: map['imageUrl']?.toString() ??
          map['image']?.toString() ??
          '',
      priceUsd: _toDouble(price?['usd'] ?? map['price']),
      priceEtb: _toDouble(price?['etb']),
      fxTimestamp: ts,
      aiMatchPercentage: _toDouble(map['aiMatchPercentage']),
      productRating: _toDouble(map['productRating'] ?? map['rating']),
      sellerScore: _toDouble(map['sellerScore']),
      taxRate: _toDouble(map['taxRate']),
      discount: _toDouble(map['discount']),
      inStock: map['inStock'] is bool ? map['inStock'] : true,
      numberSold: _toInt(map['numberSold']),
      deliveryEstimate: map['deliveryEstimate']?.toString() ?? '',
      customerHighlights: map['customerHighlights']?.toString() ?? '',
      customerReview: map['customerReview']?.toString() ?? '',
      summaryBullets: (map['summaryBullets'] is List)
          ? (map['summaryBullets'] as List)
              .map((e) => e.toString())
              .toList()
          : const [],
      deeplinkUrl: map['deeplinkUrl']?.toString() ?? '',
    );
  }

  Map<String, dynamic> toMap() => {
        'id': id,
        'title': productName,
        'description': productDescription,
        'imageUrl': imageUrl,
        'price': {
          'usd': priceUsd,
          'etb': priceEtb,
          'fxTimestamp': fxTimestamp?.toIso8601String(),
        },
        'aiMatchPercentage': aiMatchPercentage,
        'productRating': productRating,
        'sellerScore': sellerScore,
        'taxRate': taxRate,
        'discount': discount,
        'inStock': inStock,
        'numberSold': numberSold,
        'deliveryEstimate': deliveryEstimate,
        'customerHighlights': customerHighlights,
        'customerReview': customerReview,
        'summaryBullets': summaryBullets,
        'deeplinkUrl': deeplinkUrl,
      };

  dynamic operator [](String key) {
    switch (key) {
      case 'id':
        return id;
      case 'title':
      case 'name':
      case 'productName':
        return productName;
      case 'description':
      case 'productDescription':
        return productDescription;
      case 'image':
      case 'imageUrl':
        return imageUrl;
      case 'price':
      case 'priceUsd':
        return priceUsd;
      case 'priceEtb':
        return priceEtb;
      case 'fxTimestamp':
        return fxTimestamp;
      case 'aiMatchPercentage':
        return aiMatchPercentage;
      case 'productRating':
      case 'rating':
        return productRating;
      case 'sellerScore':
        return sellerScore;
      case 'taxRate':
        return taxRate;
      case 'discount':
        return discount;
      case 'inStock':
        return inStock;
      case 'numberSold':
        return numberSold;
      case 'deliveryEstimate':
        return deliveryEstimate;
      case 'customerHighlights':
        return customerHighlights;
      case 'customerReview':
        return customerReview;
      case 'summaryBullets':
        return summaryBullets;
      case 'deeplinkUrl':
        return deeplinkUrl;
      default:
        return null;
    }
  }

  ProductEntity copyWith({
    String? id,
    String? productName,
    String? productDescription,
    String? imageUrl,
    double? priceUsd,
    double? priceEtb,
    DateTime? fxTimestamp,
    double? aiMatchPercentage,
    double? productRating,
    double? sellerScore,
    double? taxRate,
    double? discount,
    bool? inStock,
    int? numberSold,
    String? deliveryEstimate,
    String? customerHighlights,
    String? customerReview,
    List<String>? summaryBullets,
    String? deeplinkUrl,
  }) {
    return ProductEntity(
      id: id ?? this.id,
      productName: productName ?? this.productName,
      productDescription: productDescription ?? this.productDescription,
      imageUrl: imageUrl ?? this.imageUrl,
      priceUsd: priceUsd ?? this.priceUsd,
      priceEtb: priceEtb ?? this.priceEtb,
      fxTimestamp: fxTimestamp ?? this.fxTimestamp,
      aiMatchPercentage: aiMatchPercentage ?? this.aiMatchPercentage,
      productRating: productRating ?? this.productRating,
      sellerScore: sellerScore ?? this.sellerScore,
      taxRate: taxRate ?? this.taxRate,
      discount: discount ?? this.discount,
      inStock: inStock ?? this.inStock,
      numberSold: numberSold ?? this.numberSold,
      deliveryEstimate: deliveryEstimate ?? this.deliveryEstimate,
      customerHighlights: customerHighlights ?? this.customerHighlights,
      customerReview: customerReview ?? this.customerReview,
      summaryBullets: summaryBullets ?? this.summaryBullets,
      deeplinkUrl: deeplinkUrl ?? this.deeplinkUrl,
    );
  }

  @override
  List<Object?> get props => [
        id,
        productName,
        productDescription,
        imageUrl,
        priceUsd,
        priceEtb,
        fxTimestamp,
        aiMatchPercentage,
        productRating,
        sellerScore,
        taxRate,
        discount,
        inStock,
        numberSold,
        deliveryEstimate,
        customerHighlights,
        customerReview,
        summaryBullets,
        deeplinkUrl,
      ];

  double get rating => productRating;
  double get price => priceUsd; // backward compatibility
  String get name => productName;

  static double _toDouble(dynamic v) {
    if (v == null) return 0;
    if (v is num) return v.toDouble();
    if (v is String) {
      final s = v.trim().replaceAll('%', '');
      return double.tryParse(s) ?? 0;
    }
    return 0;
  }

  static int _toInt(dynamic v) {
    if (v == null) return 0;
    if (v is int) return v;
    if (v is num) return v.toInt();
    if (v is String) return int.tryParse(v) ?? 0;
    return 0;
  }
}