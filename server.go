type Server struct {
}

func (s *Server) GetProfiles(ctx context.Context, req *pb.Request) (*pb.Result, error) {
	item, err := s.MemcClient.Get(i)
	if err == nil {
			// memcached hit
			// profile_str := string(item.Value)

			// fmt.Printf("memc hit\n")
			// fmt.Println(profile_str)

			hotel_prof := new(pb.Hotel)
			json.Unmarshal(item.Value, hotel_prof)
			hotels = append(hotels, hotel_prof)

	} 
}