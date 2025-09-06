
import '../../domain/Entity/comparison_table.dart';

class ComparisonTableModel extends ComparisonTable {
  const ComparisonTableModel(super.features, super.featureValuePair);

  factory ComparisonTableModel.fromJson(Map<String, dynamic> json) {

    final keys = json.keys.toList();
    final featureValuePair = json;
    print('Keys extracted: $keys'); // Debug keys
    print('featureValuePair: $featureValuePair'); // Debug featureValuePair
    return ComparisonTableModel(keys, json);
  }

  Map<String, dynamic> toJson() {
    return {'features': features, 'featureValuePair': featureValuePair};
  }

  ComparisonTable toEntity(ComparisonTableModel comparisonTable) {
    return ComparisonTable(
      comparisonTable.features,
      comparisonTable.featureValuePair,
    );
  }

  factory ComparisonTableModel.fromEntity(ComparisonTable comparisonTable) {
    return ComparisonTableModel(
      comparisonTable.features,
      comparisonTable.featureValuePair,
    );
  }
}
