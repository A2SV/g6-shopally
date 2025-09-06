import 'package:equatable/equatable.dart';
import 'package:shopallymobile/features/comparing/domain/Entity/comparison_table.dart';
import 'package:shopallymobile/features/comparing/domain/Entity/product_entity.dart';

class ComparisonEntity extends Equatable {
  final ProductEntity product;
  final List<String> pros;
  final List<String> cons;
  final bool isBest;
  final ComparisonTable table;

  const ComparisonEntity({
    required this.product,
    required this.pros,
    required this.cons,
    required this.isBest,
    required this.table,
  });

  @override
  List<Object?> get props => [product, pros, cons, isBest, table];
}
