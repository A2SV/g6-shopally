
import 'package:shopallymobile/features/comparing/data/model/over_all_comparison_model.dart';
import 'package:shopallymobile/features/comparing/domain/Entity/comparison_result_entity.dart';
import 'comparison_model.dart';

class ComparisonResultModel {
  final List<ComparisonModel> comparisonModels;
  final OverAllComparisonModel overAllComparison;

  ComparisonResultModel({
    required this.comparisonModels,
    required this.overAllComparison,
  });

  factory ComparisonResultModel.fromJson(Map<String, dynamic> json) {
    print("🟡 products key: ${json['products']}");
    print("🟡 overallComparison key: ${json['overallComparison']}");
    return ComparisonResultModel(
      comparisonModels: (json['products'] as List)
          .map((item) => ComparisonModel.fromJson(item))
          .toList(),
      overAllComparison: OverAllComparisonModel.fromJson(
        json['overallComparison'],
      ),
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'products': comparisonModels.map((e) => e.toJson()).toList(),
      'overallComparison': overAllComparison.toJson(),
    };
  }

  ComparisonResultEntity toEntity() {
    return ComparisonResultEntity(
      comparisonEntity:
      comparisonModels.map((e) => e.toEntity()).toList(),
      overAllComparisonEntity: overAllComparison.toEntity(),
    );
  }

  factory ComparisonResultModel.fromEntity(
      ComparisonResultEntity comparisonResultEntity,
      ) {
    return ComparisonResultModel(
      comparisonModels: comparisonResultEntity.comparisonEntity
          .map((e) => ComparisonModel.fromEntity(e))
          .toList(),
      overAllComparison: OverAllComparisonModel.fromEntity(
        comparisonResultEntity.overAllComparisonEntity,
      ),
    );
  }
}
