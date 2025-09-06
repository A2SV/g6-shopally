

import 'package:shopallymobile/features/comparing/domain/Entity/comparison_entity.dart';
import 'package:equatable/equatable.dart';
import 'over_all_comparison_entity.dart';

class ComparisonResultEntity extends Equatable {
  final List<ComparisonEntity> comparisonEntity;
  final OverAllComparisonEntity overAllComparisonEntity;

  const ComparisonResultEntity({
    required this.comparisonEntity,
    required this.overAllComparisonEntity,
  });

  @override
  List<Object?> get props => [comparisonEntity, overAllComparisonEntity];

}