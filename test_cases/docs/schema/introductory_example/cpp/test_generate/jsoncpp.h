#pragma once

// File automatically generated by mapry. DO NOT EDIT OR APPEND!

#include <json/json.h>  // jsoncpp

#include <map>
#include <string>

#include "parse.h"
#include "types.h"

namespace book {
namespace address {

namespace jsoncpp {

/**
 * parses Pipeline from a JSON value.
 *
 * @param [in] value to be parsed
 * @param [in] ref reference to the value (e.g., a reference path)
 * @param [out] target parsed Pipeline
 * @param [out] errors encountered during parsing
 */
void pipeline_from(
  const Json::Value& value,
  std::string ref,
  Pipeline* target,
  parse::Errors* errors);

/**
 * parses Address from a JSON value.
 *
 * @param [in] value to be parsed
 * @param ref reference to the value (e.g., a reference path)
 * @param [out] target parsed data
 * @param [out] errors encountered during parsing
 */
void address_from(
  const Json::Value& value,
  std::string ref,
  Address* target,
  parse::Errors* errors);

/**
 * parses Person from a JSON value.
 *
 * @param [in] value to be parsed
 * @param ref reference to the value (e.g., a reference path)
 * @param [out] target parsed data
 * @param [out] errors encountered during parsing
 */
void person_from(
  const Json::Value& value,
  std::string ref,
  Person* target,
  parse::Errors* errors);


/**
 * serializes Pipeline to a JSON value.
 *
 * @param pipeline to be serialized
 * @return JSON value
 */
Json::Value serialize_pipeline(
  const Pipeline& pipeline);

/**
 * serializes Person to a JSON value.
 *
 * @param person to be serialized
 * @return JSON value
 */
Json::Value serialize_person(
  const Person& person);

/**
 * serializes Address to a JSON value.
 *
 * @param address to be serialized
 * @return JSON value
 */
Json::Value serialize_address(
  const Address& address);

}  // namespace jsoncpp

}  // namespace address
}  // namespace book

// File automatically generated by mapry. DO NOT EDIT OR APPEND!