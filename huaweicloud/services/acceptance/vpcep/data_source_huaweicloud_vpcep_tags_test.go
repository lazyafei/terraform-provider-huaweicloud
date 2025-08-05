package vpcep

// import (
// 	"fmt"
// 	"regexp"
// 	"testing"

// 	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

// 	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/acceptance"
// )

// func TestAccDataTags_basic(t *testing.T) {
// 	var (
// 		all = "data.huaweicloud_vpcep_tags.test"
// 		dc  = acceptance.InitDataSourceCheck(all)
// 	)

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck: func() {
// 			acceptance.TestAccPreCheck(t)
// 		},
// 		ProviderFactories: acceptance.TestAccProviderFactories,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccDataTags_basic(),
// 				Check: resource.ComposeTestCheckFunc(
// 					dc.CheckResourceExists(),
// 					resource.TestMatchResourceAttr(all, "tags.#", regexp.MustCompile(`^[1-9]([0-9]*)?$`)),
// 					resource.TestCheckResourceAttrSet(all, "tags.0.key"),
// 					resource.TestMatchResourceAttr(all, "tags.0.values.#", regexp.MustCompile(`^[1-9]([0-9]*)?$`)),
// 					resource.TestCheckOutput("tags_validation", "true"),
// 				),
// 			},
// 		},
// 	})
// }

// func testAccDataTags_base() string {
// 	var (
// 		name = acceptance.RandomAccResourceName()
// 	)

// 	return fmt.Sprintf(`
// data "huaweicloud_er_availability_zones" "test" {}

// resource "huaweicloud_vpcep_endpoint" "test" {
//   count = 2

//   availability_zones = try(slice(data.huaweicloud_er_availability_zones.test.names, 0, 1), [])
//   name               = format("%[1]s_%%d", count.index)
//   asn                = %[2]d+count.index

//   tags = {
//     foo = format("bar%%d", count.index)
//   }

// }
// `, name)
// }

// func testAccDataTags_basic() string {
// 	return fmt.Sprintf(`
// %[1]s

// data "huaweicloud_vpcep_tags" "test" {
//   depends_on = [huaweicloud_er_instance.test]

//   resource_type = "endpoint"
// }

// output "tags_validation" {
//   value = length([for t in data.huaweicloud_er_tags.test.tags: t.key == "foo" &&
//     alltrue([for k, v in huaweicloud_er_instance.test[*].tags: contains(t.values, v) if k == "foo"])]) > 0
// }
// `, testAccDataTags_base())
// }
