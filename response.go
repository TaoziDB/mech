package youtube

type Video struct {
   DASHManifestURL string // StreamingData.DASHManifestURL
   Description string // VideoDetails.ShortDescription
   ID string
   Title string
}

type player struct {
   Microformat struct {
      PlayerMicroformatRenderer struct {
         PublishDate string
         ViewCount string
         Description struct {
            SimpleText string
         }
         Title struct {
            SimpleText string
         }
      }
   }
}