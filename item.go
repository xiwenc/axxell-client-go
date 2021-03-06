/* 
 * axxell-api
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: 1.0.0
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package AxxellClient

// An item is often equivalent to a product or service in your store
type Item struct {

	// Read-only
	CreationTime string `json:"creationTime,omitempty"`

	// This must be your product id used by your own store
	ItemId string `json:"itemId,omitempty"`

	// Global identifier of this item/product. Read-only
	Gid string `json:"gid,omitempty"`

	// Human readable title of the item/product
	Title string `json:"title,omitempty"`

	// Sanitized version of title. Read-only
	Label string `json:"label,omitempty"`

	// Categories this item belongs to. List of keywords describing the item is also acceptable here.
	Categories []string `json:"categories,omitempty"`

	// Full URL that links to the product. e.g. http://yourshop.com/product/123
	Url string `json:"url,omitempty"`
}
