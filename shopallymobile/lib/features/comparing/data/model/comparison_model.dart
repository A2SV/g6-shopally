import 'package:shopallymobile/features/comparing/data/model/product_model.dart';
import 'package:shopallymobile/features/comparing/domain/Entity/comparison_entity.dart';

import 'comparison_table_model.dart';

class ComparisonModel {
  final ProductModel productModel;
  final List<String> pros;
  final List<String> cons;
  final bool isBest;
  final ComparisonTableModel table;
  const ComparisonModel({
    required this.productModel,
    required this.pros,
    required this.cons,
    required this.isBest,
    required this.table,
  });

  factory ComparisonModel.fromJson(Map<String, dynamic> json) {


    final Map<String, dynamic> features = json['synthesis']['features'];
    features['title'] = json['product']['title'];
    print("🟡 features: $features");
    return ComparisonModel(
      productModel: ProductModel.fromJson(json['product']),
      pros: List<String>.from(json['synthesis']['pros']),
      cons: List<String>.from(json['synthesis']['cons']),
      isBest: json['synthesis']['isBestValue'],
      table: ComparisonTableModel.fromJson(features),
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'product': productModel.toJson(),
      'pros': pros,
      'cons': cons,
      'isBest': isBest,
      'table': table.toJson(),
    };
  }

  ComparisonEntity toEntity() {
    return ComparisonEntity(
      product: productModel.toEntity(),
      pros: pros,
      cons: cons,
      isBest: isBest,
      table: table.toEntity(table),
    );
  }

  factory ComparisonModel.fromEntity(ComparisonEntity comparison) {
    return ComparisonModel(
      productModel: ProductModel.fromEntity(comparison.product),
      pros: comparison.pros,
      cons: comparison.cons,
      isBest: comparison.isBest,
      table: ComparisonTableModel.fromEntity(comparison.table),
    );
  }
}
