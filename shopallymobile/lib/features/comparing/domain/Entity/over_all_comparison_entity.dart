
import 'package:equatable/equatable.dart';
import 'package:shopallymobile/features/comparing/domain/Entity/price_entity.dart';

class OverAllComparisonEntity extends Equatable{

  final String bestValueProduct;
  final String bestValueLink;
  final PriceEntity bestValuePrice;
  final List<String> keyHighlights;
  final String summary;

  const OverAllComparisonEntity({
    required this.bestValueProduct,
    required this.bestValueLink,
    required this.bestValuePrice,
    required this.keyHighlights,
    required this.summary,
  });

  @override
  List<Object?> get props => [bestValueProduct, bestValueLink, bestValuePrice, keyHighlights, summary];


}