import '../../domain/entities/product_entity.dart';

class ProductModel extends ProductEntity {
  final double? etbPrice;

  const ProductModel({
    required String productName,
    required String productDescription,
    required String id,
    required double priceUsd,
    required String imageUrl,
    required double productRating,
    bool inStock = true,
    double aiMatchPercentage = 0.0,
    double? etbPrice,
    DateTime? fxTimestamp,
    double sellerScore = 0.0,
    String deliveryEstimate = '',
    String customerHighlights = '',
    String customerReview = '',
    int numberSold = 0,
    List<String> summaryBullets = const [],
    String deeplinkUrl = '',
    double taxRate = 0.0,
    double discount = 0.0,
  })  : etbPrice = etbPrice,
        super(
          productName: productName,
          productDescription: productDescription,
          id: id,
          priceUsd: priceUsd,
          imageUrl: imageUrl,
          productRating: productRating,
          inStock: inStock,
          aiMatchPercentage: aiMatchPercentage,
          fxTimestamp: fxTimestamp,
          sellerScore: sellerScore,
          deliveryEstimate: deliveryEstimate,
          customerHighlights: customerHighlights,
          customerReview: customerReview,
          numberSold: numberSold,
          summaryBullets: summaryBullets,
          deeplinkUrl: deeplinkUrl,
          taxRate: taxRate,
          discount: discount,
        );

  factory ProductModel.fromJson(Map<String, dynamic> json) {
    final priceObj = json['price'] as Map<String, dynamic>?;

    return ProductModel(
      id: json['id'] as String,
      productName: json['title'] as String,
      productDescription: json['description'] as String? ?? '',
      imageUrl: json['imageUrl'] as String? ?? '',
      priceUsd: (priceObj?['usd'] as num?)?.toDouble() ?? 0.0,
      productRating: (json['productRating'] as num?)?.toDouble() ?? 0.0,
      aiMatchPercentage: (json['aiMatchPercentage'] as num?)?.toDouble() ?? 0.0,
      etbPrice: (priceObj?['etb'] as num?)?.toDouble(),
      fxTimestamp: priceObj?['fxTimestamp'] != null
          ? DateTime.tryParse(priceObj!['fxTimestamp'] as String)
          : null,
      sellerScore: (json['sellerScore'] as num?)?.toDouble() ?? 0.0,
      deliveryEstimate: (json['deliveryEstimate'] as String?) ?? '',
      customerHighlights: (json['customerHighlights'] as String?) ?? '',
      customerReview: (json['customerReview'] as String?) ?? '',
      numberSold: (json['numberSold'] as num?)?.toInt() ?? 0,
      summaryBullets: (json['summaryBullets'] as List<dynamic>?)
              ?.map((e) => e.toString())
              .toList() ??
          const [],
      deeplinkUrl: (json['deeplinkUrl'] as String?) ?? '',
      taxRate: (json['taxRate'] as num?)?.toDouble() ?? 0.0,
      discount: (json['discount'] as num?)?.toDouble() ?? 0.0,
      inStock: json['inStock'] as bool? ?? true,
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'title': productName,
      'description': productDescription,
      'imageUrl': imageUrl,
      'price': {
        'usd': priceUsd,
        'etb': etbPrice,
        'fxTimestamp': fxTimestamp?.toIso8601String(),
      },
      'productRating': productRating,
      'aiMatchPercentage': aiMatchPercentage,
      'sellerScore': sellerScore,
      'deliveryEstimate': deliveryEstimate,
      'customerHighlights': customerHighlights,
      'customerReview': customerReview,
      'numberSold': numberSold,
      'summaryBullets': summaryBullets,
      'deeplinkUrl': deeplinkUrl,
      'taxRate': taxRate,
      'discount': discount,
      'inStock': inStock,
    };
  }

  @override
  List<Object?> get props => [
        productName,
        productDescription,
        id,
        priceUsd,
        imageUrl,
        inStock,
        productRating,
        aiMatchPercentage,
        etbPrice,
        fxTimestamp,
        sellerScore,
        deliveryEstimate,
        customerHighlights,
        customerReview,
        numberSold,
        summaryBullets,
        deeplinkUrl,
        taxRate,
        discount,
      ];
}