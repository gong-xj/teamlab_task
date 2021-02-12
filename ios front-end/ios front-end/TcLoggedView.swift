
import SwiftUI

struct TcLoggedView: View {
    @State var id: String
    @State var name: String
    @State var stData: [St]
    @State var stOrTe: String
    @State var vercode: String
    
    var body: some View {
        TabView {
            StListView(stData: stData, vercode:vercode)
                .tabItem {
                     Image(systemName: "list.bullet")
                     Text("成績")
                }
            
            MeView(name:name, id:id, stOrTe:stOrTe)
                .tabItem {
                    Image(systemName: "person.fill")
                    Text("個人情報")
                }
        }
    }
}

struct TcLoggedView_Previews: PreviewProvider {
    static var previews: some View {
        TcLoggedView(id: "id??", name: "name??", stData: [], stOrTe: "先生？？", vercode: "vercode??")
    }
}
