import 'package:equatable/equatable.dart';

class ComparisonTable extends Equatable {
  final List<String> features;
  final Map<String, dynamic> featureValuePair;

  const ComparisonTable(this.features, this.featureValuePair);

  @override
  List<Object?> get props => [features, featureValuePair];
}
