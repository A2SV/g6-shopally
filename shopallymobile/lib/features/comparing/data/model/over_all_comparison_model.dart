import 'package:shopallymobile/features/comparing/data/model/price_model.dart';
import 'package:shopallymobile/features/comparing/domain/Entity/over_all_comparison_entity.dart';


class OverAllComparisonModel {
  final String bestValueProduct;
  final String bestValueLink;
  final PriceModel bestValuePrice;
  final List<String> keyHighlights;
  final String summary;

  const OverAllComparisonModel({
    required this.bestValueProduct,
    required this.bestValueLink,
    required this.bestValuePrice,
    required this.keyHighlights,
    required this.summary,
  });

  factory OverAllComparisonModel.fromJson(Map<String, dynamic> json) {
    return OverAllComparisonModel(
      bestValueProduct: json['bestValueProduct'],
      bestValueLink: json['bestValueLink'],
      bestValuePrice: PriceModel.fromJson(json['bestValuePrice']),
      keyHighlights: List<String>.from(json['keyHighlights']),
      summary: json['summary'],
    );
  }

  Map<String, dynamic> toJson() => {
    'bestValueProduct': bestValueProduct,
    'bestValueLink': bestValueLink,
    'bestValuePrice': bestValuePrice.toJson(),
    'keyHighlights': keyHighlights,
    'summary': summary,
  };

  OverAllComparisonEntity toEntity() {
    return OverAllComparisonEntity(
      bestValueProduct: bestValueProduct,
      bestValueLink: bestValueLink,
      bestValuePrice: bestValuePrice.toEntity(bestValuePrice),
      keyHighlights: keyHighlights,
      summary: summary,
    );
  }

  factory OverAllComparisonModel.fromEntity(
      OverAllComparisonEntity overAllComparisonEntity,
      ) {
    return OverAllComparisonModel(
      bestValueProduct: overAllComparisonEntity.bestValueProduct,
      bestValueLink: overAllComparisonEntity.bestValueLink,
      bestValuePrice: PriceModel.fromEntity(
        overAllComparisonEntity.bestValuePrice,
      ),
      keyHighlights: overAllComparisonEntity.keyHighlights,
      summary: overAllComparisonEntity.summary,
    );
  }
}
