#pragma once

// File automatically generated by mapry. DO NOT EDIT OR APPEND!

#include <map>
#include <memory>
#include <string>

namespace some {
namespace graph {

struct SomeGraph;

class Empty;

// defines an empty class.
class Empty {
public:
  // identifies the instance.
  std::string id;
};

// defines some object graph.
struct SomeGraph {
  // tests a map of class references.
  std::map<std::string, Empty*> map_of_class_refs;

  // registers Empty instances.
  std::map<std::string, std::unique_ptr<Empty>> empties;
};

}  // namespace graph
}  // namespace some

// File automatically generated by mapry. DO NOT EDIT OR APPEND!