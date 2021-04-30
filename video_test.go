package youtube
import "testing"

var tests = []struct {
   id, desc string
   dash bool
} {
   {
      "XeojXq6ySs4",
      "Provided to YouTube by Epitaph\n\nSnowflake · Kate Bush\n\n" +
      "50 Words For Snow\n\n" +
      "℗ Noble & Brite Ltd. trading as Fish People, under exclusive license to Anti Inc.\n\n" +
      "Released on: 2011-11-22\n\nMusic  Publisher: Noble and Brite Ltd.\n" +
      "Composer  Lyricist: Kate Bush\n\nAuto-generated by YouTube.",
      false,
   }, {
      "ClYg-0-z_ds",
      "",
      false,
   }, {
      "BnEn7X3Pr7o",
      "",
      true,
   },
}

func TestVideo(t *testing.T) {
   for _, test := range tests {
      v, err := NewVideo(test.id)
      if err != nil {
         t.Error(err)
      }
      if v.Description() != test.desc {
         t.Error(v)
      }
      if test.dash && v.StreamingData.DashManifestURL == "" {
         t.Errorf("%+v", v)
      }
   }
}
